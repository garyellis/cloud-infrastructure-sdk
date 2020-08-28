package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const vmwareVariablesTf = "variables.tf"

type VMwareVariablesTf struct {
	input.Input
	EnvName string
	AppName string
}

func (t *VMwareVariablesTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, vmwareTfModuleDir, vmwareVariablesTf)
	}
	t.TemplateBody = vmwareVariablesTfTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const vmwareVariablesTfTmpl = `
variable "name" {
  description = "a unique label to identify all resources"
  type = string
}

variable "vsphere_dc_name" {
  type = string
}

variable "vsphere_cluster_name" {
  type = string
}

variable "vsphere_datastore_name" {
  type = string
}

variable "vsphere_network_names" {
  type = list(string)
}

variable "folder" {
  type = string
}

variable "domain" {
  type = string
  default = "ews.int"
}

variable "ipv4_gateway" {
  type = string
}

variable "vsphere_vm_template_name" {
  description = "the vmware template name"
  type        = string
  default     = "TurboOps-Cent7-cloud-init"
}

variable "nodes" {
  description = "A list of hostname and ip addresses"
  type = list(map(string))
  default = []
}

variable "nodes_num_cpus" {
  type = string
  default = "2"
}

variable "nodes_memory" {
  type = string
  default = "4096"
}

variable "nodes_additional_disks" {
  description = "A list of additional disks attached to the nodes"
  type        = list(map(string))
  default     = []
}

#### terraform ssh provisioner configuration (runs on creation only)
variable "bootstrap_remote_exec" {
  type = list(string)
  default = ["echo foo"]
}

variable "provisioner_ssh_user" {
  type = string
  default = "provisioner"
}

variable "provisioner_ssh_password" {
  type = string
  default = ""
}

variable "vault_addr" {
  description = "The vault address"
  type = string
}

variable "vault_ssh_ca_path" {
  description = "The vault ssh ca mount path"
  type        = string
  default     = "ssh-client-signer"
}

variable "vault_secret_kv_path" {
  description = "The vault kv path used by cloud-init userdata"
  type        = string
  default     = "secret/edap/os-common"
}

variable "vault_secret_vas_username_key" {
  description = "The vault secret username key name"
  type        = string
  default     = "username"
}

variable "vault_secret_vas_password_key" {
  description = "The vault secret password key name"
  type        = string
  default     = "ad_secret"
}
`
