// Modified from github.com/kubernetes-sigs/controller-tools/pkg/scaffold/input/input.go

package input

import "text/template"

type IfExistsAction int

const (

	// Overwrite truncates and overwrites the existing file (by default)
	Overwrite IfExistsAction = iota

	// Error returns an error and stops processing
	Error

	// Skip skips the file and moves onto the next
	Skip
)

// Input is the input for scaffolding a file
type Input struct {
	// Path is the file to be written
	Path string

	//
	IfExistsAction IfExistsAction
	// IsExec indicates that the file should be written with executable permission
	IsExec bool

	// TemplateBody is the template to execute
	TemplateBody string

	// TemplateFuncs is additional functions used in the template.
	// these must be registered before execution
	TemplateFuncs template.FuncMap

	// Repo is the project repo name
	Repo string

	// ProjectName Is the project name if different from the repo name
	ProjectName string

	// AbsProjectPath is the absolute path of the project parent folder
	AbsProjectPath string
}

type Repo interface {
	SetRepo(string)
}

func (i *Input) SetRepo(r string) {
	if i.Repo == "" {
		i.Repo = r
	}
}

type ProjectName interface {
	SetProjectName(string)
}

func (i *Input) SetProjectName(n string) {
	if i.ProjectName == "" {
		i.ProjectName = n
	}
}

type AbsProjectPath interface {
	SetAbsProjectPath(string)
}

func (i *Input) SetAbsProjectPath(p string) {
	if i.AbsProjectPath == "" {
		i.AbsProjectPath = p
	}
}

// File is a scaffoldale file
type File interface {
	GetInput() (Input, error)
}

// Validate validates the input
type Validate interface {
	// Validate returns nil if the inputs' validation logic allows
	// field values, or the template.
	Validate() error
}

// Config configures the execution scaffold templates
type Config struct {
	Repo           string
	AbsProjectPath string
	ProjectName    string
}
