package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Scripts          `yaml:"scripts,omitempty" json:"scripts,omitempty"`
	AnsibleTerraform `yaml:"ansible_terraform,omitempty" json:"ansible_terraform,omitempty"`
}

type Scripts struct {
	TerraformHelpers `yaml:"terraform_helpers,omitempty"`
	VaultHelpers     `yaml:"vault_helpers,omitempty"`
	DockerHelpers    `yaml:"docker_helpers,omitempty"`
}

type DockerHelpers struct {
	DockerCeVersion        string `yaml:"docker_ce_version,omitempty"`
	DockerCeURL            string `yaml:"docker_ce_url,omitempty"`
	CiPipelinesDockerImage string `yaml:"pipeline_docker_image,omitempty"`
}

type TerraformHelpers struct {
	TerraformVersion      string `yaml:"terraform_version,omitempty"`
	TerragruntVersion     string `yaml:"terragrunt_version,omitempty"`
	TerraformReleasesURL  string `yaml:"terraform_releases_url,omitempty"`
	TerragruntDownloadURL string `yaml:"terragrunt_download_url,omitempty"`
	TfenvRepoURL          string `yaml:"tfenv_repo_url,omitempty"`
}

type VaultHelpers struct {
	VaultSshCertPrincipal    string `yaml:"vault_ssh_cert_principal,omitempty"`
	VaultSshClientSignerPath string `yaml:"vault_ssh_client_signer_path,omitempty"`
}

func NewConfig() *Config {
	config := &Config{
		Scripts{
			TerraformHelpers{
				TerraformVersion:      "0.12.26",
				TerraformReleasesURL:  "https://releases.hashicorp.com/terraform/",
				TerragruntVersion:     "v0.24.0",
				TerragruntDownloadURL: "https://github.com/gruntwork-io/terragrunt/releases/download/$TERRAGRUNT_VERSION/terragrunt_linux_amd64",
				TfenvRepoURL:          "https://github.com/tfutils/tfenv.git",
			},
			VaultHelpers{
				VaultSshCertPrincipal:    "",
				VaultSshClientSignerPath: "",
			},
			DockerHelpers{
				DockerCeVersion:        "",
				DockerCeURL:            "",
				CiPipelinesDockerImage: "",
			},
		},
		AnsibleTerraform{
			TerraformModuleSources: TerraformModuleSources{
				CloudInit:             "github.com/garyellis/cloud-init",
				VaultApprole:          "github.com/garyellis/vault-approle",
				SecurityGroup:         "github.com/garyellis/security-group",
				Ec2Instance:           "github.com/garyellis/ec2-instance",
				NetworkLoadBalancer:   "github.com/garyellis/network-loadbalancer",
				Route53Zone:           "github.com/garyellis/route53-zone",
				VsphereVirtualMachine: "github.com/garyellis/vsphere-virtualmachine",
			},
		},
	}

	return config
}

func (config *Config) ReadConfigFile(path string) error {
	yamlfile, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlfile, config)
	return err
}
