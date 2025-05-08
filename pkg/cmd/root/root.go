package root

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/cli"
	"github.com/wallanaq/oidc-cli/pkg/cmd/fetch"
)

func NewRootCommand(opts *cli.Options) *cobra.Command {

	rootCmd := &cobra.Command{
		Use:          "oidc",
		Short:        "OIDC command-line tool",
		Long:         "oidc-cli is a tool for performing OIDC operations.",
		Version:      GetVersion(),
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	globalFlags := &opts.GlobalFlags

	flags := rootCmd.PersistentFlags()

	flags.BoolVarP(&globalFlags.Verbose, "verbose", "v", false, "enable verbose output")
	flags.StringVarP(&globalFlags.Output, "output", "o", "-", "write file(s) to directory, instead of STDOUT")

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.SetVersionTemplate(VersionTemplate())

	rootCmd.AddGroup(&cobra.Group{
		ID:    "oidc",
		Title: "OIDC Commands:",
	})

	rootCmd.AddCommand(fetch.NewFetchCommand(opts))

	return rootCmd

}
