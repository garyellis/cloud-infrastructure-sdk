package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntAwsHclFile = "terragrunt.hcl"

type TerragruntAwsHcl struct {
	input.Input
	EnvName string
	AppName string
	DCName  string
}

func (t *TerragruntAwsHcl) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfLiveBaseDir, t.DCName, t.EnvName, t.AppName, terragruntAwsHclFile)
	}
	t.TemplateBody = terragruntAwsHclTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const terragruntAwsHclTmpl = `

`
