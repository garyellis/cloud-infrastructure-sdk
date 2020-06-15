// Modified from github.com/operator-framework/operator-sdk/blob/master/internal/util/fileutil/file_util.go

package fileutil

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/afero"
)

const (
	DefaultDirFileMode  = 0750
	DefaultFileMode     = 0644
	DefaultExecFileMode = 0755

	DefaultFileFlags = os.O_WRONLY | os.O_CREATE
)

// FileWriter is a file creation wrapper
type FileWriter struct {
	fs   afero.Fs
	once sync.Once
}

func NewFileWriterFS(fs afero.Fs) *FileWriter {
	fw := &FileWriter{}
	fw.once.Do(func() {
		fw.fs = fs
	})
	return fw
}

func (fw *FileWriter) GetFS() afero.Fs {
	fw.once.Do(func() {
		fw.fs = afero.NewOsFs()
	})
	return fw.fs
}

func (fw *FileWriter) WriteCloser(path string, mode os.FileMode) (io.Writer, error) {
	dir := filepath.Dir(path)
	err := fw.GetFS().MkdirAll(dir, DefaultDirFileMode)
	if err != nil {
		return nil, err
	}

	fi, err := fw.GetFS().OpenFile(path, DefaultFileFlags, mode)
	if err != nil {
		return nil, err
	}

	return fi, nil

}

func (fw *FileWriter) WriteFile(filePath string, content []byte) error {
	f, err := fw.WriteCloser(filePath, DefaultFileMode)
	if err != nil {
		return fmt.Errorf("failed to create %s: %v", filePath, err)
	}

	if c, ok := f.(io.Closer); ok {
		defer func() {
			if err := c.Close(); err != nil {
				log.Fatal(err)
			}
		}()
	}

	_, err = f.Write(content)
	if err != nil {
		return fmt.Errorf("failed to write %s: %v", filePath, err)
	}
	return nil
}

func IsClosedError(e error) bool {
	pathErr, ok := e.(*os.PathError)
	if !ok {
		return false
	}
	if pathErr.Err == os.ErrClosed {
		return true
	}
	return false
}

func GetCwd() string {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get working directory: (%v)", err)
	}
	return wd
}
