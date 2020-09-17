package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntAwsVarsFile = "vars.yaml"

type TerragruntAwsVars struct {
	input.Input
	EnvName   string
	AppName   string
	DCName    string
	AWSRegion string
	VaultAddr string
}

func (t *TerragruntAwsVars) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfLiveBaseDir, t.DCName, t.EnvName, terragruntAwsVarsFile)
	}
	t.TemplateBody = terragruntAwsVarsTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const terragruntAwsVarsTmpl = `---
name: {{.AppName}}-{{.DCName}}-{{.EnvName}}


region: "{{.AWSRegion}}"
allowed_account_ids: []
tags:
  dcname: {{.DCName}}
  environment: {{.EnvName}}

{{.AppName}}:
  nodes_count: 1
  nodes_instance_type: t3.medium
  ami_id: ami-3ecc8f46
  key_name: ""
  disable_api_termination: false
  instance_auto_recovery_enabled: false
  vpc_id: ""
  nodes_subnet_ids: []
  sg_attachments: []
  sg_egress_cidr_rules: []
  sg_ingress_cidr_rules: []
  lb_subnet_ids: []

dns_domain: ""
dns_zone_id: ""

vault_addr: "{{.VaultAddr}}"
vault_ssh_ca_path: ""
`
