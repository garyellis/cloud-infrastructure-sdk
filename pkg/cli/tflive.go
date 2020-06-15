package cli

import (
	sdkcmd "github.com/garyellis/cloud-infrastructure-sdk/pkg/cmd"
	"github.com/spf13/cobra"
)

// TerraformLiveCmd prints the version
func TerraformLiveCmd() *cobra.Command {
	tfLiveCmd := &cobra.Command{
		Use:   "terraform-live",
		Short: "create terraform live projects",
	}
	tfLiveCmd.PersistentFlags().StringVar(&projectName, "project-name", "terraform-live", "the teraform live project name")
	//tfLiveCmd.PersistentFlags().StringVar(&envName, "env-name", "development", "the terraform live environment name")
	tfLiveCmd.PersistentFlags().StringVar(&appName, "app-name", "app-name", "the terraform live live subfolder name")
	tfLiveCmd.PersistentFlags().StringSliceVarP(&envNames, "env-name", "e", []string{"development"}, "one or more live environment names")
	tfLiveCmd.AddCommand(InitProjectCmd())
	return tfLiveCmd
}

func InitProjectCmd() *cobra.Command {
	initProjectCmd := &cobra.Command{
		Use:   "init",
		Short: "creates a new terraform live project",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := sdkcmd.DoTfLiveScaffold(projectName, envNames, appName)
			return err
		},
	}
	return initProjectCmd
}
