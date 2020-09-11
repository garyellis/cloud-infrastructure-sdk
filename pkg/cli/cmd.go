package cli

import (
	"github.com/spf13/cobra"
)

const cliName = "cloud-infra-sdk"

var (
	projectName   string
	envName       string
	envNames      []string
	appName       string
	dcName        string
	infraProvider string
	configFile    string
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     cliName,
		Short:   "coud infrastructure sdk",
		Version: Version,
	}

	cmd.AddCommand(VersionCmd())
	cmd.AddCommand(TerraformLiveCmd())
	cmd.AddCommand(AnsibleTerraformCmd())
	return cmd
}
