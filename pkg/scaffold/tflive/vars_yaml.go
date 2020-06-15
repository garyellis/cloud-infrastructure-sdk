package tflive

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const varsYamlFile = "vars.yaml"

type VarsYaml struct {
	input.Input
	EnvName string
	AppName string
}

func (y *VarsYaml) GetInput() (input.Input, error) {
	if y.Path == "" {
		y.Path = filepath.Join(LiveBaseDir, y.EnvName, varsYamlFile)
	}
	y.TemplateBody = VarsYamlTmpl

	y.IfExistsAction = input.Skip
	return y.Input, nil
}

const VarsYamlTmpl = `---
name: {{.AppName}}

tags:
  environment-stage: {{.EnvName}}
`
