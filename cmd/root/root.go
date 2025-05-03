package root

import (
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/cmd/fetch"
	"github.com/wallanaq/oidc-cli/internal/cli"
)

const (
	versionTemplate = "oidc-cli/{{.Version}}\n"
)

func NewRootCommand(opts *cli.Options) *cobra.Command {

	rootCmd := &cobra.Command{
		Use:          "oidc",
		Short:        "OIDC command-line tool",
		Long:         "oidc-cli is a tool for performing OIDC operations.",
		Version:      "v0.1.0",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if opts.GlobalFlags.Verbose {
				opts.Logger.SetOutput(os.Stderr)
				opts.Logger.Println("Verbose logging enabled")
			} else {
				opts.Logger.SetOutput(io.Discard)
			}
		},
	}

	flags := rootCmd.PersistentFlags()

	flags.BoolVarP(&opts.GlobalFlags.Verbose, "verbose", "v", false, "enable verbose output")
	flags.StringVarP(&opts.GlobalFlags.Output, "output", "o", "-", "write file(s) to directory, instead of STDOUT")

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.SetVersionTemplate(versionTemplate)

	rootCmd.AddGroup(&cobra.Group{
		ID:    "oidc",
		Title: "OIDC Commands:",
	})

	rootCmd.AddCommand(fetch.NewFetchCommand(opts))

	return rootCmd

}
