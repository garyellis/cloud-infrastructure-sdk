package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntVMwareVarsFile = "vars.yaml"

type VsphereNode struct {
	Name        string `yaml:"name,omitempty"`
	Hostname    string `yaml:"hostname,omitempty"`
	IPV4Address string `yaml:"ipv4_address,omitempty"`
}
type TerragruntVMwareVars struct {
	input.Input
	EnvName                string
	AppName                string
	DCName                 string
	VaultAddr              string
	VaultSshCAPath         string
	VsphereDCName          string              `yaml:"vsphere_dc_name,omitempty"`
	VsphereClusterName     string              `yaml:"vsphere_cluster_name,omitempty"`
	VsphereFolder          string              `yaml:"vsphere_folder,omitempty"`
	VsphereIpv4Gateway     string              `yaml:"vsphere_ipv4_gateway,omitempty"`
	VsphereNetworkNames    []string            `yaml:"vsphere_network_names,omitempty"`
	VsphereDataStoreName   string              `yaml:"vsphere_datastore_name,omitempty"`
	VsphereNodes           []VsphereNode       `yaml:"vsphere_nodes,omitempty"`
	VsphereAdditionalDisks []map[string]string `yaml:"vsphere_additional_disks,omitempty"`
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

vsphere_dc_name: "{{.VsphereDCName}}"
vsphere_cluster_name: "{{.VsphereClusterName}}"
vsphere_folder: "{{.VsphereFolder}}"
vsphere_datastore_name: "{{.VsphereDataStoreName}}"
vsphere_ipv4_gateway: "{{.VsphereIpv4Gateway}}"
vsphere_network_names: [{{ range $index, $network := .VsphereNetworkNames }}{{if $index}},{{end}}"{{ $network }}"{{ end }}]

{{.AppName}}:
  num_cpus: 2
  memory: 8192
  additional_disks: []
{{- if .VsphereNodes }}
  nodes:{{ range .VsphereNodes }}
    - { name: "{{.Name}}", hostname: "{{.Hostname}}", ipv4_address: "{{.IPV4Address}}" }
{{- end}}
{{ else }}
  nodes: []
{{- end }}

vault_addr: "{{.VaultAddr}}"
vault_ssh_ca_path: "{{.VaultSshCAPath}}"

vault_secret_vas_username_key: ""
vault_secret_vas_password_key: ""
`
