package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const awsOutputsTf = "outputs.tf"

type AWSOutputsTf struct {
	input.Input
	AppName string
}

func (t *AWSOutputsTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, awsTfModuleDir, awsOutputsTf)
	}
	t.TemplateBody = awsOutputsTfTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const awsOutputsTfTmpl = `
locals {
  ansible_hostvars = [
    for i, host in compact(concat(module.nodes.aws_instance_private_ips)) : {
  	  "host"               = host
      vault_wrapping_token = element(concat(module.vault_approle.wrapping_tokens, list("")), i)
    }
  ]
}

output "ansible_inventory" {
  value = templatefile("${path.module}/ansible_inventory.yml.tmpl", {
    lb_fqdn = local.lb_fqdn
    vault_addr = var.vault_addr
    vault_role_id = module.vault_approle.role_id
    vault_secret_id_response_wrapping_path = module.vault_approle.secret_id_response_wrapping_path
    vault_pki_role_name = module.vault_approle.pki_secret_backend_role_name
    hostvars = local.ansible_hostvars
    nodes = module.nodes.aws_instance_private_ips,
  })
  sensitive = true
}

output "vault_role_id" {
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
