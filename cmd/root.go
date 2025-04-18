package cmd

import (
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/cmd/discovery"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:          "oidc",
	Version:      "v0.1.0",
	Short:        "OIDC command-line tool",
	Long:         "oidc-cli is a tool for performing OIDC operations",
	SilenceUsage: true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if verbose {
			log.SetOutput(os.Stdout)
		} else {
			log.SetOutput(io.Discard)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

	rootCmd.AddGroup(discovery.DiscoveryGroup)
	rootCmd.AddCommand(discovery.DiscoveryCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
