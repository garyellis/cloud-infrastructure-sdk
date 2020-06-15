package tflive

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntHclFile = "terragrunt.hcl"

type TerragruntHcl struct {
	input.Input
	EnvName string
	AppName string
}

func (t *TerragruntHcl) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(LiveBaseDir, t.EnvName, t.AppName, terragruntHclFile)
	}
	t.TemplateBody = TerragruntHclTmpl

	return t.Input, nil
}

const TerragruntHclTmpl = `include {
  path = find_in_parent_folders()	
}

terraform source {
	source = "../../../modules/{{.AppName}}/aws"
}

locals {
	version = yamldecode(file("${get_terragrunt_dir()}/${find_in_parent_folders("../metadata.yaml")}"))
	vars    = yamldecode(file("${get_terragrunt_dir()}/${find_in_parent_folders("vars.yaml")}"))
  }

  inputs = {
## user defined inputs
  }
`
