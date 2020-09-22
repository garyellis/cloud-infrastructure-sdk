package config

import (
	"io/ioutil"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/ansibleterraform"
	"gopkg.in/yaml.v2"
)

type AnsibleTerraform struct {
	S3BucketRegion         string `yaml:"s3_bucket_region,omitempty"`
	S3BucketNamePrefix     string `yaml:"s3_bucket_name_prefix,omitempty"`
	TerraformModuleSources `yaml:"terraform_modules,omitempty"`
	AnsibleRoleSources     `yaml:"ansible_roles,omitempty"`
}

type TerraformModuleSources struct {
	CloudInit             string `yaml:"cloud_init,omitempty"`
	VaultApprole          string `yaml:"vault_approle,omitempty"`
	SecurityGroup         string `yaml:"security_group,omitempty"`
	Ec2Instance           string `yaml:"ec2_instance,omitempty"`
	NetworkLoadBalancer   string `yaml:"network_loadbalancer,omitempty"`
	Route53Zone           string `yaml:"route53_zone,omitempty"`
	VsphereVirtualMachine string `yaml:"vsphere_virtualmachine,omitempty"`
}

type AnsibleRoleSources struct {
	AppRoleSources []ansibleterraform.AnsibleRole `yaml:"app_roles,omitempty"`
	OSRoleSources  []ansibleterraform.AnsibleRole `yaml:"os_roles,omitempty"`
}

type TerragruntVarsConfig ansibleterraform.TerragruntAwsVars

func NewTerragruntVarsConfig() *TerragruntVarsConfig {
	c := &TerragruntVarsConfig{}
	return c
}

func (c *TerragruntVarsConfig) ReadConfigFile(path string) error {
	yamlfile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlfile, c)
	return err
}
