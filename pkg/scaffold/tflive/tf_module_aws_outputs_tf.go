package tflive

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const tfModuleAwsOutputsTfFile = "outputs.tf"

type TfModuleAwsOutputsTf struct {
	input.Input
	AppName string
}

func (t *TfModuleAwsOutputsTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(ModuleBaseDir, t.AppName, "aws", tfModuleAwsOutputsTfFile)
	}
	t.TemplateBody = tfModuleAwsOutputsTfTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const tfModuleAwsOutputsTfTmpl = `#### terraform aws infrastructure module outputs
`
