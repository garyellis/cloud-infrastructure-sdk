package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const osYmlFile = "os.yml"

type OSYml struct {
	input.Input
	OSRoles []AnsibleRole
}

func (t *OSYml) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsiblePlaybooksDir, osYmlFile)
	}
	t.TemplateBody = osYmlTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const osYmlTmpl = `---
- name: converge the os roles
  hosts: all
  roles:{{range $element := .OSRoles}}
    - {{.Name}}
  {{- end}}
`
