package version

import (
	"runtime"

	"github.com/spf13/cobra"
)

var version string

const versionTemplate = "oidc-cli version %s (%s/%s)\n"

func NewVersionCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of oidc-cli",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf(versionTemplate, version, runtime.GOOS, runtime.GOARCH)
		},
	}

}
