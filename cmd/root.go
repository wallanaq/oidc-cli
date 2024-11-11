package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/cmd/auth"
)

var Verbose bool
var Version string = "v0.1.0"

var RootCmd = &cobra.Command{
	Use:   "oidc",
	Short: "CLI for OIDC authentication",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if Verbose {
			log.SetOutput(os.Stdout)
		} else {
			log.SetOutput(io.Discard)
		}
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	RootCmd.AddCommand(auth.LoginCmd)
	RootCmd.AddCommand(auth.LogoutCmd)
}
