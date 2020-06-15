package tflive

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const tfModuleAwsVariablesTfFile = "variables.tf"

type TfModuleAwsVariablesTf struct {
	input.Input
	AppName string
}

func (t *TfModuleAwsVariablesTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(ModuleBaseDir, t.AppName, "aws", tfModuleAwsVariablesTfFile)
	}
	t.TemplateBody = tfModuleAwsMainTfTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const tfModuleAwsVariablesTfTmpl = `#### terraform aws infrastructure module variables
`
