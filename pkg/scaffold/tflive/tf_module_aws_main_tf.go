package tflive

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const tfModuleAwsMainTfFile = "main.tf"

type TfModuleAwsMainTf struct {
	input.Input
	AppName string
}

func (t *TfModuleAwsMainTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(ModuleBaseDir, t.AppName, "aws", tfModuleAwsMainTfFile)
	}
	t.TemplateBody = tfModuleAwsMainTfTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const tfModuleAwsMainTfTmpl = `#### terraform aws infrastructure module contents
`
