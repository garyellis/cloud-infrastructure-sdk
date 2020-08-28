package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const awsVariablesTf = "variables.tf"

type AWSVariablesTf struct {
	input.Input
	EnvName string
	AppName string
}

func (t *AWSVariablesTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, awsTfModuleDir, awsVariablesTf)
	}
	t.TemplateBody = awsVariablesTfTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const awsVariablesTfTmpl = `
variable "allowed_account_ids" {
  description = "the aws provider allowed account ids"
  type = list(string)
}

variable "region" {
  description = "the aws provider region"
  type        = string
  }

variable "name" {
  description = "a unique label to identify all resources"
  type = string
}

variable "dns_domain" {
  description = "the route53 dns domain name"
  type        = string
}

variable "dns_zone_id" {
  description = "the route53 private zone id"
  type        = string
}
  
variable "ami_id" {
  description = "The jump server AMI ID"
  type        = string
}
  
variable "key_name" {
  description = "The jump server keypair name"
  type        = string
  default     = ""
}

variable "iam_instance_profile" {
  description = "an iam instance profile arn attached to all nodes"
  type        = string
  default     = ""
}

variable "disable_api_termination" {
  description = "Enable ec2 instance termination protection"
  type        = bool
  default     = false
}

variable "instance_auto_recovery_enabled" {
  description = "Enable ec2 instance autorecovery"
  type        = bool
  default     = false
}

variable "vpc_id" {
  description = "the target vpc id"
  type = string
}

variable "sg_attachments" {
  description = "security group attachments on ec2 instances"
  type        = list(string)
  default     = []
}

variable "sg_egress_cidr_rules" {
  description = "security group egress cidr rules"
  type        = list(map(string))
  default     = []
}
  
variable "sg_ingress_cidr_rules" {
  description = "security group ingress cidr rules"
  type        = list(map(string))
  default     = []
}

#### nodes
variable "nodes_count" {
  description = "The number of nodes"
  type        = number
  default     = 2
}

variable "nodes_instance_type" {
  description = "The nodes instance type"
  type        = string
  default     = "t3.large"
  }

variable "nodes_subnet_ids" {
  description = "the ec2 instances subnet ids"
  type = list(string)
  default = []
}

variable "lb_subnet_ids" {
  description = "The nlb subnet ids"
  type = list(string)
  default = []
}

variable "tags" {
  description = "A map of tags on all taggable resources"
  type        = map(string)
  default     = {}
}

variable "vault_addr" {
  description = "The vault address"
  type        = string
}

variable "vault_ssh_ca_path" {
  description = "The vault ssh ca mount path"
  type        = string
  default     = "ssh-client-signer"
}
`
