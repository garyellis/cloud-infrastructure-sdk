package cli

import (
	sdkcmd "github.com/garyellis/cloud-infrastructure-sdk/pkg/cmd"
	"github.com/spf13/cobra"
)

// AnsibleTerraformCmd prints the version
func AnsibleTerraformCmd() *cobra.Command {
	ansibleTerraformCmd := &cobra.Command{
		Use:   "ansible-terraform",
		Short: "create an ansible-terraform project",
	}
	ansibleTerraformCmd.PersistentFlags().StringVar(&configFile, "config", "config.yaml", "the teraform live project name")
	ansibleTerraformCmd.PersistentFlags().StringVar(&projectName, "project-name", "my-project", "the teraform live project name")
	ansibleTerraformCmd.PersistentFlags().StringVar(&appName, "app-name", "my-app", "the terraform live live subfolder name")
	ansibleTerraformCmd.PersistentFlags().StringVar(&infraProvider, "infra-provider", "aws", "infrastructure provider. Valid providers are aws and vmware")
	ansibleTerraformCmd.PersistentFlags().StringVar(&dcName, "dc-name", "my-dc", "the data center name")
	ansibleTerraformCmd.PersistentFlags().StringSliceVarP(&envNames, "env-name", "e", []string{"development"}, "one or more environment names")
	ansibleTerraformCmd.PersistentFlags().StringVar(&vaultAddr, "vault-addr", "https://vault-demo.ews.works", "The hashicorp vault server")
	ansibleTerraformCmd.PersistentFlags().StringVar(&vaultSSHCa, "vault-ssh-ca", "", "The hashicorp vault ssh ca secret engine path")
	ansibleTerraformCmd.PersistentFlags().StringVar(&vaultSSHRole, "vault-ssh-role", "", "The hashicorp vault ssh role name")
	ansibleTerraformCmd.PersistentFlags().StringVar(&sshUser, "ssh-user", "$USER", "The ssh username")
	ansibleTerraformCmd.PersistentFlags().StringVar(&awsRegion, "aws-region", "us-west-2", "The aws region when the infrastructure provider is type aws")
	ansibleTerraformCmd.PersistentFlags().StringVar(&s3BucketName, "s3-bucket-name", "ews-works", "The remote state s3 bucket name")
	ansibleTerraformCmd.PersistentFlags().StringVar(&s3BucketRegion, "s3-bucket-region", "us-west-2", "The remote state s3 bucket region")
	ansibleTerraformCmd.PersistentFlags().StringVar(&vSphereServer, "vsphere-server", "", "The vCenter server name for vSphere API operations")
	ansibleTerraformCmd.PersistentFlags().StringVar(&terragruntVarsFile, "terragrunt-vars", "terragrunt-vars.yaml", "the terragrunt vars file")
	ansibleTerraformCmd.AddCommand(InitAnsibleTerraformProjectCmd())
	return ansibleTerraformCmd
}

func InitAnsibleTerraformProjectCmd() *cobra.Command {
	initProjectCmd := &cobra.Command{
		Use:   "init",
		Short: "creates a new ansible-terraform project",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := sdkcmd.InitAnsibleTerraformScaffold(configFile, terragruntVarsFile, cliName, Version, projectName, appName, infraProvider, dcName, envNames, vaultAddr, vaultSSHCa, vaultSSHRole, sshUser, awsRegion, s3BucketName, s3BucketRegion, vSphereServer)
			return err
		},
	}
	return initProjectCmd
}
