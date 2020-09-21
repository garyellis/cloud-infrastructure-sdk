package scripts

import (
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terraformPluginsTxtFile = "terraform-plugins"

type TerraformPluginsTxt struct {
	input.Input
	PluginUrls []string
}

func (t *TerraformPluginsTxt) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = terraformPluginsTxtFile
	}
	t.TemplateBody = terraformPluginsTxtTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const terraformPluginsTxtTmpl = `{{- range $url := .PluginUrls -}}
{{ $url }}
{{ end -}}`
