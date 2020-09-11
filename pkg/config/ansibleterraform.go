package config

type AnsibleTerraform struct {
	S3BucketRegion         string `yaml:"s3_bucket_region,omitempty"`
	S3BucketNamePrefix     string `yaml:"s3_bucket_name_prefix,omitempty"`
	TerraformModuleSources `yaml:"terraform_modules,omitempty"`
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
