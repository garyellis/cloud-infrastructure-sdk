package tflive

import (
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const readmeMdFile = "README.md"

type ReadmeMd struct {
	input.Input
}

func (r *ReadmeMd) GetInput() (input.Input, error) {
	if r.Path == "" {
		r.Path = readmeMdFile
	}
	r.TemplateBody = ReadmeMdTmpl

	return r.Input, nil
}

const ReadmeMdTmpl = `# terraform live infrastructure for {{.ProjectName}}
`
