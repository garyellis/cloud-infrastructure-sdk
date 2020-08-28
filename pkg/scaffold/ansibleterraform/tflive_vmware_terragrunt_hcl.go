package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntVMwareHclFile = "terragrunt.hcl"

type TerragruntVMwareHcl struct {
	input.Input
	EnvName string
	AppName string
	DCName  string
}

func (t *TerragruntVMwareHcl) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfLiveBaseDir, t.DCName, t.EnvName, t.AppName, terragruntVMwareHclFile)
	}
	t.TemplateBody = terragruntVMwareHclTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const terragruntVMwareHclTmpl = `include {
  path = find_in_parent_folders()
}

terraform {
  source = "../../../../modules/{{.AppName}}/vmware"
}

locals {
  vars    = yamldecode(file("${get_terragrunt_dir()}/${find_in_parent_folders("vars.yaml")}"))
}

inputs = {
  name                             = local.vars.name
  vsphere_dc_name                  = local.vars.vsphere_dc_name
  vsphere_cluster_name             = local.vars.vsphere_cluster_name
  vsphere_datastore_name           = local.vars.vsphere_datastore_name
  vsphere_network_names            = local.vars.vsphere_network_names
  folder                           = local.vars.vsphere_folder
  ipv4_gateway                     = local.vars.vsphere_ipv4_gateway
  nodes                            = local.vars.{{.AppName}}.nodes
  nodes_num_cpus                   = local.vars.{{.AppName}}.num_cpus
  nodes_memory                     = local.vars.{{.AppName}}.memory
  nodes_additional_disks           = local.vars.{{.AppName}}.additional_disks

  vault_addr                        = local.vars.vault_addr
  vault_ssh_ca_path                 = local.vars.vault_ssh_ca_path
  vault_secret_vas_username_key     = local.vars.vault_secret_vas_username_key
  vault_secret_vas_password_key     = local.vars.vault_secret_vas_password_key
}
`
