package cmd

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var verbose bool

var rootCmd = &cobra.Command{
	Use:     "oidc",
	Version: "v0.1.0",
	Short:   "A OIDC command-line tool",
	Long:    "oidc-cli is a tool for performing OIDC operations",
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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
