package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terragruntAwsVarsFile = "vars.yaml"

type TerragruntAwsVars struct {
	input.Input
	EnvName        string
	AppName        string
	DCName         string
	AWSRegion      string
	VaultAddr      string
	VaultSshCAPath string
	AwsAccountID   string   `yaml:"aws_account_id,omitempty"`
	AmiID          string   `yaml:"ami_id,omitempty"`
	VpcID          string   `yaml:"vpc_id,omitempty"`
	LBSubnetIDs    []string `yaml:"lb_subnet_ids,omitempty"`
	NodesSubnetIDs []string `yaml:"nodes_subnet_ids,omitempty"`
	DNSDomain      string   `yaml:"dns_domain,omitempty"`
	DNSZoneID      string   `yaml:"dns_zone_id,omitempty"`
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
allowed_account_ids:
  - "{{.AwsAccountID}}"
tags:
  dcname: {{.DCName}}
  environment: {{.EnvName}}

{{.AppName}}:
  nodes_count: 1
  nodes_instance_type: t3.medium
  ami_id: {{ default "ami-3ecc8f46" .AmiID}}
  key_name: ""
  disable_api_termination: false
  instance_auto_recovery_enabled: false
  vpc_id: "{{.VpcID}}"
  nodes_subnet_ids: [{{ range $index, $subnet := .NodesSubnetIDs }}{{if $index}},{{end}}"{{ $subnet }}"{{ end }}]
  sg_attachments: []
  sg_egress_cidr_rules: []
  sg_ingress_cidr_rules: []
  lb_subnet_ids: [{{ range $index, $subnet := .LBSubnetIDs }}{{if $index}},{{end}}"{{ $subnet }}"{{ end }}]

dns_domain: "{{.DNSDomain}}"
dns_zone_id: "{{.DNSZoneID}}"

vault_addr: "{{.VaultAddr}}"
vault_ssh_ca_path: "{{.VaultSshCAPath}}"
`
