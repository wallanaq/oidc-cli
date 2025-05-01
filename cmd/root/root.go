package root

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/cmd/fetch"
	"github.com/wallanaq/oidc-cli/internal/cli"
)

func NewRootCommand(opts *cli.Options) *cobra.Command {

	cmd := &cobra.Command{
		Use:          "oidc",
		Short:        "OIDC command-line tool",
		Long:         "oidc-cli is a tool for performing OIDC operations.",
		Version:      "v0.1.0",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if opts.Flags.Verbose {
				opts.Logger.SetOutput(os.Stderr)
				opts.Logger.Println("Verbose logging enabled")
			} else {
				opts.Logger.SetOutput(io.Discard)
			}
		},
	}

	flags := cmd.PersistentFlags()

	flags.BoolVarP(&opts.Flags.Verbose, "verbose", "v", false, "enable verbose logging")
	flags.StringVarP(&opts.Flags.Output, "output", "o", "-", "output file path (default: stdout)")

	cmd.AddGroup(&cobra.Group{ID: "oidc", Title: "OIDC Commands:"})

	cmd.AddCommand(
		fetch.NewFetchCommand(opts),
	)

	return cmd

}
