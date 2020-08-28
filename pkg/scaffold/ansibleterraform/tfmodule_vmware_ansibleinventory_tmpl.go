package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const vmwareAnsibleInventoryYml = "ansible_inventory.yml.tmpl"

type VMwareAnsileInventory struct {
	input.Input
	AppName string
}

func (t *VMwareAnsileInventory) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, vmwareTfModuleDir, vmwareAnsibleInventoryYml)
	}
	t.TemplateBody = vmwareAnsibleInventoryTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const vmwareAnsibleInventoryTmpl = `---
all:
  hosts:
%{ for i in hostvars ~}
    ${i.host}:
      vault_wrapping_token: ${i.vault_wrapping_token}
%{ endfor ~}
  vars:
    vault_addr: "${vault_addr}"
    vault_role_id: "${vault_role_id}"
    vault_secret_id_response_wrapping_path: "${vault_secret_id_response_wrapping_path}"
    vault_pki_role_name: "${vault_pki_role_name}"
  children:
  {{.AppName}}:
      vars:
      hosts:
%{ for ip in ibm_sp_nodes ~}
        ${ip}:
%{ endfor ~}
`
