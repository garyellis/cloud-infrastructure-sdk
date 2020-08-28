package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntAwsVarsFile = "vars.yml"

type TerragruntAwsVars struct {
	input.Input
	EnvName string
	AppName string
	DCName  string
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


region: ""
allowed_account_ids: []
tags:
  dcname: {{.DCName}}
  environment: {{.EnvName}}

{{.AppName}}:
  nodes_count: 1
  nodes_instance_type:
  ami_id:
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

vault_addr: ""
vault_ssh_ca_path: ""
`
