package root

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/cmd/fetch"
	"github.com/wallanaq/oidc-cli/internal/cli"
)

var Version = "v0.0.0"

func NewRootCommand(opts *cli.Options) *cobra.Command {

	rootCmd := &cobra.Command{
		Use:          "oidc",
		Short:        "OIDC command-line tool",
		Long:         "oidc-cli is a tool for performing OIDC operations.",
		Version:      Version,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger := opts.Logger
			if opts.GlobalFlags.Verbose {
				logger.SetOutput(os.Stderr)
				logger.Println("Verbose logging enabled")
			} else {
				logger.SetOutput(io.Discard)
			}
		},
	}

	flags := rootCmd.PersistentFlags()

	globalFlags := &opts.GlobalFlags

	flags.BoolVarP(&globalFlags.Verbose, "verbose", "v", false, "enable verbose output")
	flags.StringVarP(&globalFlags.Output, "output", "o", "-", "write file(s) to directory, instead of STDOUT")

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.SetVersionTemplate("oidc-cli/{{.Version}}\n")

	rootCmd.AddGroup(&cobra.Group{
		ID:    "oidc",
		Title: "OIDC Commands:",
	})

	rootCmd.AddCommand(fetch.NewFetchCommand(opts))

	return rootCmd

}
