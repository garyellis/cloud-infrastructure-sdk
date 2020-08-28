package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntVMwareVarsFile = "vars.yml"

type TerragruntVMwareVars struct {
	input.Input
	EnvName string
	AppName string
	DCName  string
}

func (t *TerragruntVMwareVars) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfLiveBaseDir, t.DCName, t.EnvName, terragruntVMwareVarsFile)
	}
	t.TemplateBody = terragruntVMwareVarsTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const terragruntVMwareVarsTmpl = `---
name: {{.AppName}}-{{.DCName}}-{{.EnvName}}

vsphere_dc_name:
vsphere_cluster_name:
vsphere_folder:
vsphere_datastore_name:
vsphere_ipv4_gateway:
vsphere_network_names:

{{.AppName}}:
  num_cpus:
  memory:
  additional_disks: []
  nodes: []

vault_addr: ""
vault_ssh_ca_path: ""
vault_secret_kv_path: ""
vault_secret_vas_username_key: ""
vault_secret_vas_password_key: ""

`
