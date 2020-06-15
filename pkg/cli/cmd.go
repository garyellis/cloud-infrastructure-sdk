package cli

import (
	"github.com/spf13/cobra"
)

const cliName = "cloud-infra-sdk"

var (
	projectName string
	envName     string
	envNames    []string
	appName     string
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     cliName,
		Short:   "coud infrastructure sdk",
		Version: Version,
	}

	cmd.AddCommand(VersionCmd())
	cmd.AddCommand(TerraformLiveCmd())
	return cmd
}
