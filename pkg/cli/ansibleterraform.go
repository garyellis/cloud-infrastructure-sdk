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
	ansibleTerraformCmd.PersistentFlags().StringVar(&projectName, "project-name", "my-project", "the teraform live project name")
	ansibleTerraformCmd.PersistentFlags().StringVar(&appName, "app-name", "my-app", "the terraform live live subfolder name")
	ansibleTerraformCmd.PersistentFlags().StringVar(&infraProvider, "infra-provider", "aws", "infrastructure provider. Valid providers are aws and vmware")
	ansibleTerraformCmd.PersistentFlags().StringVar(&dcName, "dc-name", "my-dc", "the data center name")
	ansibleTerraformCmd.PersistentFlags().StringSliceVarP(&envNames, "env-name", "e", []string{"development"}, "one or more environment names")
	ansibleTerraformCmd.AddCommand(InitAnsibleTerraformProjectCmd())
	return ansibleTerraformCmd
}

func InitAnsibleTerraformProjectCmd() *cobra.Command {
	initProjectCmd := &cobra.Command{
		Use:   "init",
		Short: "creates a new ansible-terraform project",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := sdkcmd.InitAnsibleTerraformScaffold(cliName, Version, projectName, appName, infraProvider, dcName, envNames)
			return err
		},
	}
	return initProjectCmd
}
