package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const requirementsYmlFile = "requirements.yml"

type AnsibleRole struct {
	Src     string
	Name    string
	Version string
}

type RequirementsYml struct {
	input.Input
	AppRoles []AnsibleRole
	OSRoles  []AnsibleRole
}

func (t *RequirementsYml) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsibleBaseDir, requirementsYmlFile)
	}
	t.TemplateBody = requirementsYmlTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const requirementsYmlTmpl = `---
# os roles{{range $element := .OSRoles }}
- name: {{.Name}}
  src: {{.Src}}
  version: {{.Version}}
{{- end}}
# application roles{{range $element := .AppRoles }}
- name: {{.Name}}
  src: {{.Src}}
  version: {{.Version}}
{{- end}}
`
