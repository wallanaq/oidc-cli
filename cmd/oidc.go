package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var cmdFlags struct {
	Verbose bool
}

// NewOIDCCommand creates the `oidc` command and its nested children.
func NewOIDCCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "oidc",
		Short:        "OIDC command-line tool",
		Long:         "oidc-cli is a tool for performing OIDC operations.",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if cmdFlags.Verbose {
				log.SetOutput(os.Stderr)
				log.Println("Verbose logging enabled")
			} else {
				log.SetOutput(io.Discard)
			}
		},
	}

	cmd.Flags().BoolVarP(&cmdFlags.Verbose, "verbose", "v", false, "enable verbose logging")

	return cmd
}
