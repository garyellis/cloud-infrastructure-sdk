// Modified from github.com/kubernetes-sigs/controller-tools/pkg/scaffold/scaffold.go

package scaffold

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/util/fileutil"
	"github.com/spf13/afero"
)

// Scaffold writes Templates to scaffold new files
type Scaffold struct {
	CliName string

	CliVersion string

	// Repo is the project repo name
	Repo string

	AbsProjectPath string

	ProjectName string

	// Fs is the filesystem that GetWriter uses to write scaffold files
	Fs afero.Fs

	// GetWriter returns a writer for writing scaffold files.
	GetWriter func(path string, mode os.FileMode) (io.Writer, error)
}

func (s *Scaffold) setFieldsAndValidate(t input.File) error {
	if c, ok := t.(input.CliName); ok {
		c.SetCliName(s.CliName)
	}

	if c, ok := t.(input.CliVersion); ok {
		c.SetCliVersion(s.CliVersion)
	}

	if r, ok := t.(input.Repo); ok {
		r.SetRepo(s.Repo)
	}
	if p, ok := t.(input.ProjectName); ok {
		p.SetProjectName(s.ProjectName)
	}
	if a, ok := t.(input.AbsProjectPath); ok {
		a.SetAbsProjectPath(s.AbsProjectPath)
	}

	if v, ok := t.(input.Validate); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

func (s *Scaffold) configure(cfg *input.Config) {
	s.CliName = cfg.CliName
	s.CliVersion = cfg.CliVersion
	s.Repo = cfg.Repo
	s.AbsProjectPath = cfg.AbsProjectPath
	s.ProjectName = cfg.ProjectName
}

// Execute renders the input scaffolding
func (s *Scaffold) Execute(cfg *input.Config, files ...input.File) error {
	if s.Fs == nil {
		s.Fs = afero.NewOsFs()
	}
	if s.GetWriter == nil {
		s.GetWriter = fileutil.NewFileWriterFS(s.Fs).WriteCloser
	}

	s.configure(cfg)

	for _, f := range files {
		if err := s.doFile(f); err != nil {
			return err
		}
	}
	return nil
}

func (s *Scaffold) doFile(e input.File) error {
	err := s.setFieldsAndValidate(e)
	if err != nil {
		return err
	}

	i, err := e.GetInput()
	if err != nil {
		return err
	}

	absFilePath := filepath.Join(s.AbsProjectPath, i.Path)

	if _, err := s.Fs.Stat(absFilePath); err == nil || os.IsExist(err) {
		switch i.IfExistsAction {
		case input.Overwrite:
			//log.Println("File exists. Overwriting.", i.Path)
		case input.Skip:
			//log.Println("File exists. Skipping", i.Path)
			return nil
		case input.Error:
			return fmt.Errorf("%s already exists", absFilePath)
		}
	}

	return s.doRender(i, e, absFilePath)
}
func (s *Scaffold) doRender(i input.Input, e input.File, absPath string) error {
	var mode os.FileMode = fileutil.DefaultFileMode
	if i.IsExec {
		mode = fileutil.DefaultExecFileMode
	}
	f, err := s.GetWriter(absPath, mode)
	if err != nil {
		return err
	}
	if c, ok := f.(io.Closer); ok {
		defer func() {
			if err := c.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}

	var b []byte
	if c, ok := e.(CustomRenderer); ok {
		c.SetFs(s.Fs)
		// CustomRenderers have a non template method for file rendering
		if _, err := c.CustomRender(); err != nil {
			return err
		}
	} else {
		tmpl, err := newTemplate(i)
		if err != nil {
			return err
		}

		out := &bytes.Buffer{}
		if err = tmpl.Execute(out, e); err != nil {
			return err
		}
		b = out.Bytes()
	}

	//
	if _, err = s.Fs.Stat(absPath); err == nil && i.IfExistsAction == input.Overwrite {
		if file, ok := f.(afero.File); ok {
			if err = file.Truncate(0); err != nil {
				return err
			}
		}
	}

	_, err = f.Write(b)
	log.Print("Created ", i.Path)
	return err
}

func newTemplate(i input.Input) (*template.Template, error) {
	t := template.New(i.Path).Funcs(template.FuncMap{
		"title": strings.Title,
		"lower": strings.ToLower,
	})
	if len(i.TemplateFuncs) > 0 {
		t.Funcs(i.TemplateFuncs)
	}
	return t.Parse(i.TemplateBody)
}

type CustomRenderer interface {
	SetFs(afero.Fs)
	CustomRender() ([]byte, error)
}
