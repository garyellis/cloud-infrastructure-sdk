package ansibleterraform

import (
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const readmeMdFile = "README.md"

type ReadmeMd struct {
	input.Input
}

func (t *ReadmeMd) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = readmeMdFile
	}
	t.TemplateBody = readmeMdTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const readmeMdTmpl = `# ansible/terraform live infrastructure for {{.ProjectName}}
`
