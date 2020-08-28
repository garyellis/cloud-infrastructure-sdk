package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const vmwareMainTf = "main.tf"
const vmwareTfModuleDir = "vmware"

type VMwareMainTf struct {
	input.Input
	EnvName string
	AppName string
}

func (t *VMwareMainTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, vmwareTfModuleDir, vmwareMainTf)
	}
	t.TemplateBody = vmwareMainTfTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const vmwareMainTfTmpl = `
locals {
  vault_count_approle_secret_wrapping_tokens = length(var.nodes) 
}

module "vault_approle" {
  source = "git::ssh://git@stash.ews.int:7999/terrm/vault-auth.git?ref=v0.2.1"

  name                                = var.name
  vault_count_approle_wrapping_tokens = local.vault_count_approle_secret_wrapping_tokens
  vault_kv_paths                      = []
}


module "cloud_init_secret" {
  source = "git::ssh://git@stash.ews.int:7999/terrm/vault-wrapping-token.git?ref=v0.1.0"

  count_wrapping_tokens = length(var.nodes)
  ttl                   = 900
  path                  = var.vault_secret_kv_path
}

locals {
  user_data_per_instance = [ for i, _ in var.nodes:
    base64gzip(templatefile("${path.module}/userdata.sh.tmpl", {
      vault_addr                  = var.vault_addr
      vault_ssh_ca_path           = var.vault_ssh_ca_path
      kv_secret_wrapping_token    = module.cloud_init_secret.wrapping_tokens[i]
      kv_secret_vas_username_key  = var.vault_secret_vas_username_key
      kv_secrret_vas_password_key = var.vault_secret_vas_password_key
    }))
  ]
}

module "nodes" {
  source = "git::ssh://git@stash.ews.int:7999/terrm/ews-vsphere-virtual-machine.git?ref=v0.3.1"

  vsphere_dc_name          = var.vsphere_dc_name
  vsphere_cluster_name     = var.vsphere_cluster_name
  vsphere_datastore_name   = var.vsphere_datastore_name
  vsphere_network_names    = var.vsphere_network_names
  ipv4_gateway             = var.ipv4_gateway
  create_folder            = false
  folder                   = var.folder
  num_cpus                 = var.nodes_num_cpus
  memory                   = var.nodes_memory
  additional_disks         = var.nodes_additional_disks
  virtualmachines          = var.nodes
  vsphere_vm_template_name = var.vsphere_vm_template_name
  user_data_per_instance   = local.user_data_per_instance

  provisioner_ssh_user     = var.provisioner_ssh_user
  provisioner_ssh_password = var.provisioner_ssh_password
}
`
