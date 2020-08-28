package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const vmwareOutputsTf = "outputs.tf"

type VMwareOutputsTf struct {
	input.Input
	AppName string
}

func (t *VMwareOutputsTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, vmwareTfModuleDir, vmwareOutputsTf)
	}
	t.TemplateBody = vmwareOutputsTfTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const vmwareOutputsTfTmpl = `
locals {
  ansible_hostvars = [
    for i, host in compact(concat(
      module.nodes.virtualmachines.*.ipv4_address,
    )) : {
      "host"               = host
      vault_wrapping_token = element(concat(module.vault_approle.wrapping_tokens, list("")), i)
    }
  ]
}

output "ansible_inventory" {
  value = templatefile("${path.module}/ansible_inventory.yml.tmpl", {
    vault_addr                             = var.vault_addr
    vault_role_id                          = module.vault_approle.role_id
    vault_secret_id_response_wrapping_path = module.vault_approle.secret_id_response_wrapping_path
    vault_pki_role_name                    = module.vault_approle.pki_secret_backend_role_name
    hostvars                               = local.ansible_hostvars
    nodes                        = module.nodes.virtualmachines.*.ipv4_address
  })
  sensitive = true
}

output "vault_approle_id" {
  value = module.vault_approle.role_id
}

output "vault_secret_id_response_wrapping_path" {
  value = module.vault_approle.secret_id_response_wrapping_path
}

output "wrapping_tokens" {
  value     = module.vault_approle.wrapping_tokens
  sensitive = true
}
`
