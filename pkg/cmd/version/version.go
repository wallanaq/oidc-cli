package version

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	version     string
	buildNumber string
)

const versionTemplate string = "oidc-cli version %s-%s (%s/%s)"

func NewVersionCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "version",
		Short: "Print the version of oidc-cli",
		Run: func(cmd *cobra.Command, args []string) {
			v := fmt.Sprintf(versionTemplate, version, buildNumber, runtime.GOOS, runtime.GOARCH)
			cmd.Println(v)
		},
	}

}
