package tflive

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const tfModuleAwsLocalsTfFile = "locals.tf"

type TfModuleAwsLocalsTf struct {
	input.Input
	AppName string
}

func (t *TfModuleAwsLocalsTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(ModuleBaseDir, t.AppName, "aws", tfModuleAwsLocalsTfFile)
	}
	t.TemplateBody = tfModuleAwsLocalsTfTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const tfModuleAwsLocalsTfTmpl = `#### terraform aws infrastructure local variables
`
