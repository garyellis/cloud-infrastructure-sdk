package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const awsLocalsTf = "locals.tf"

type AWSLocalsTf struct {
	input.Input
	AppName string
}

func (t *AWSLocalsTf) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, awsTfModuleDir, awsLocalsTf)
	}
	t.TemplateBody = awsLocalsTfTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const awsLocalsTfTmpl = `
data "aws_subnet" "lb_subnets" {
	for_each = toset(var.lb_subnet_ids)
	id       = each.value
  }
  
locals {
  
  #### static security group rules
  # computed "static" variables
  lb_subnet_cidrs       = [for i in data.aws_subnet.lb_subnets : i.cidr_block]

  # {{.AppName}} static rules
  rules = [
    { desc = "{{.AppName}} example rule ", from_port = "8443", to_port = "8443", protocol = "tcp" },
  ]
  ingress_sg_rules = []
  ingress_cidr_rules = [
    { desc = "nlb {{.AppName}} example", from_port = "8443", to_port = "8443", protocol = "tcp", cidr_blocks = join(",", local.lb_subnet_cidrs) },
  ]
  
  #### lb configuration
  lb_listeners                  = [{ port = "443", target_group_index = "0" },]
  lb_listeners_count            = 1
  lb_target_groups              = [{ name = "8443", target_type = "ip", port = "8443",},]
  lb_target_groups_count        = 1
  lb_target_group_health_checks = [
    { target_groups_index = "0", protocol = "TCP", port = "8443", interval = "10", healthy_threshold = "2", unhealthy_threshold = "2" },
  ]

  }
`
