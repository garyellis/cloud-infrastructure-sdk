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

const terragruntAwsHclTmpl = `include {
  path = find_in_parent_folders()
}

terraform {
  source = "../../../../modules/{{.AppName}}/aws"
}

locals {
  vars = yamldecode(file("${get_terragrunt_dir()}/../vars.yaml"))
}

inputs = {
  allowed_account_ids             = local.vars.allowed_account_ids
  region                          = local.vars.region
  name                            = local.vars.name
  tags                            = local.vars.tags
  dns_domain                      = local.vars.dns_domain
  dns_zone_id                     = local.vars.dns_zone_id
  vpc_id                          = local.vars.{{.AppName}}.vpc_id
  vault_addr                      = local.vars.vault_addr
  vault_ssh_ca_path               = local.vars.vault_ssh_ca_path

  nodes_count                     = local.vars.{{.AppName}}.nodes_count
  nodes_instance_type             = local.vars.{{.AppName}}.nodes_instance_type
  ami_id                          = local.vars.{{.AppName}}.ami_id
  key_name                        = local.vars.{{.AppName}}.key_name
  disable_api_termination         = local.vars.{{.AppName}}.disable_api_termination
  instance_auto_recovery_enabled  = local.vars.{{.AppName}}.instance_auto_recovery_enabled
  sg_attachments                  = local.vars.{{.AppName}}.sg_attachments
  sg_egress_cidr_rules            = local.vars.{{.AppName}}.sg_egress_cidr_rules
  sg_ingress_cidr_rules           = local.vars.{{.AppName}}.sg_ingress_cidr_rules
  nodes_subnet_ids                = local.vars.{{.AppName}}.nodes_subnet_ids
  lb_subnet_ids                   = local.vars.{{.AppName}}.lb_subnet_ids
}
`
