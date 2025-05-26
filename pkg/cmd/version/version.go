package version

import (
	"runtime"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/version"
)

func NewVersionCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of oidc-cli",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf(version.Template, version.Current, runtime.GOOS, runtime.GOARCH)
		},
	}

}
