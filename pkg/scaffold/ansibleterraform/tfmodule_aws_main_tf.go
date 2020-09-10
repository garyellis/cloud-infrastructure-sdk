package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const awsMainTf = "main.tf"
const awsTfModuleDir = "aws"

type AWSMainTf struct {
	input.Input
	EnvName                           string
	AppName                           string
	TFModuleVaultApproleSource        string
	TFModuleCloudInitSource           string
	TFModuleEc2InstanceSource         string
	TFModuleSecurityGroupSource       string
	TFModuleNetworkLoadBalancerSource string
	TFModuleRoute53ZoneSource         string
}

func (t *AWSMainTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, awsTfModuleDir, awsMainTf)
	}
	t.TemplateBody = awsMainTfTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const awsMainTfTmpl = `
provider "aws" {
  allowed_account_ids = var.allowed_account_ids
  region              = var.region
}

module "cloud_init" {
  source = "{{.TFModuleCloudInitSource}}"

  base64_encode          = false
  gzip                   = false
  extra_user_data_script = templatefile("${path.module}/userdata.sh.tmpl", {
    vault_addr = var.vault_addr
    vault_ssh_ca_path = var.vault_ssh_ca_path
  })
  }

module "vault_approle" {
  source = "{{.TFModuleVaultApproleSource}}"

  name                                = var.name
  vault_count_approle_wrapping_tokens = var.nodes_count
  vault_kv_paths                      = []
  vault_pki_role_allowed_domains      = list(var.dns_domain)
}

module "sg" {
  source = "{{.TFModuleSecurityGroupSource}}"

  description                      = var.name
  self_security_group_rules        = local.rules
  ingress_security_group_rules     = local.ingress_sg_rules
  egress_cidr_rules                = var.sg_egress_cidr_rules
  ingress_cidr_rules               = concat(var.sg_ingress_cidr_rules, local.ingress_cidr_rules)
  name                             = var.name
  tags                             = var.tags
  vpc_id                           = var.vpc_id
}

module "nodes" {
  source = "{{.TFModuleEc2InstanceSource}}"

  count_instances                = var.nodes_count
  disable_api_termination        = var.disable_api_termination
  instance_auto_recovery_enabled = var.instance_auto_recovery_enabled
  ami_id                         = var.ami_id
  iam_instance_profile           = var.iam_instance_profile
  associate_public_ip_address    = false
  instance_type                  = var.nodes_instance_type
  root_block_device              = [{
    delete_on_termination = "true",
     encrypted = "true"
  }]
  ebs_block_device               = var.ebs_block_device

  key_name                       = var.key_name
  name                           = var.name
  security_group_attachments     = concat(list(module.sg.security_group_id), var.sg_attachments)
  subnet_ids                     = var.nodes_subnet_ids
  tags                           = var.tags
  user_data                      = module.cloud_init.cloudinit_userdata
}

module "lb" {
  source = "{{.TFModuleNetworkLoadBalancerSource}}"

  name                       = var.name
  enable_deletion_protection = false
  internal                   = true
  listeners_count            = local.lb_listeners_count
  listeners                  = local.lb_listeners
  subnets                    = var.lb_subnet_ids
  target_groups_count        = local.lb_target_groups_count
  target_groups              = local.lb_target_groups
  target_group_health_checks = local.lb_target_group_health_checks
  vpc_id                     = var.vpc_id
  tags                       = var.tags
}

resource "aws_lb_target_group_attachment" "lb" {
  count                      = var.nodes_count

  target_group_arn           = module.lb.target_group_arns[0]
  target_id                  = module.nodes.aws_instance_private_ips[count.index]
}

#### load balancers dns
locals {
  lb_name = var.name
  lb_fqdn = format("%s.%s", local.lb_name, var.dns_domain)
}

module "lb_dns" {
  source = "{{.TFModuleRoute53ZoneSource}}"

  create_zone           = false
  name                  = var.dns_domain
  alias_records         = [
    { name = local.lb_name, aws_dns_name = module.lb.lb_dns_name, zone_id = module.lb.lb_zone_id, evaluate_target_health = "true" },
  ]
  alias_records_count   = 1
  zone_id               = var.dns_zone_id
}
`
