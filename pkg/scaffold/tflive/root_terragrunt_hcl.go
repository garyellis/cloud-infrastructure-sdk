package tflive

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const rootTerragruntHclFile = "terragrunt.hcl"

type RootTerragruntHcl struct {
	input.Input
	EnvName string
	AppName string
}

func (t *RootTerragruntHcl) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(LiveBaseDir, t.EnvName, rootTerragruntHclFile)
	}
	t.TemplateBody = RootTerragruntHclTmpl

	return t.Input, nil
}

const RootTerragruntHclTmpl = `generate "partial_s3_backend" {
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
		bucket = "${BUCKET_NAME}"
		key = "cloud-infra-sdk/terraform/{{.ProjectName}}/live/{{.EnvName}}"
		region = "us-west-2"
		encrypt = true
	}
}
`
