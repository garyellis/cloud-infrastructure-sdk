package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const middlewareYmlFile = "middleware.yml"

type MiddlewareYml struct {
	input.Input
	AppName string
}

func (t *MiddlewareYml) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsiblePlaybooksDir, middlewareYmlFile)
	}
	t.TemplateBody = middlewareYmlTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const middlewareYmlTmpl = `---
- name: converge the application roles
  hosts: {{.AppName}}
  roles:{{range $element := .OSRoles}}
    - {{.Name}}
  {{- end}}
`
