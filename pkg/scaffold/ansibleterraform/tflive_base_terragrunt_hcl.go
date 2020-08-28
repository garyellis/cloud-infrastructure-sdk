package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntBaseHclFile = "terragrunt.hcl"

type TerragruntBaseHcl struct {
	input.Input
	EnvName        string
	DCName         string
	S3BucketName   string
	S3BucketRegion string
	S3KeyPrefix    string
}

func (t *TerragruntBaseHcl) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfLiveBaseDir, t.DCName, t.EnvName, terragruntBaseHclFile)
	}
	t.TemplateBody = terragruntBaseHclTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const terragruntBaseHclTmpl = `generate "partial_s3_backend" {
  path = "terraform.tf"
  if_exists = "overwrite"
  contents = <<EOF
terraform {
  backend "s3" {}
}
EOF
}

remote_state {
  backend = "s3
    config = {
      bucket = "{{.S3BucketName}}"
      key = "cloud-infra-sdk/{{.ProjectName}}/iaas/terraform/live/{{.DCName}}/{{.EnvName}}/${path_relative_to_include()/terraform.tfstate"
      region = "{{.S3BucketRegion}}"
      encrypt = true
    }
}
`
