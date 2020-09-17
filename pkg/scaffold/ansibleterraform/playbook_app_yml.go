package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const applicationYmlFile = "app.yml"

type ApplicationYml struct {
	input.Input
	AppName      string
	AnsibleRoles []AnsibleRole
}

func (t *ApplicationYml) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsiblePlaybooksDir, applicationYmlFile)
	}
	t.TemplateBody = applicationYmlTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const applicationYmlTmpl = `---
- name: converge the application roles
  hosts: {{.AppName}}
  roles:{{range $element := .AnsibleRoles}}
    - {{.Name}}
  {{- end}}
`
