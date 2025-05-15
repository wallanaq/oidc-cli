package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/config"
	"github.com/wallanaq/oidc-cli/pkg/cmd/root"
)

func main() {

	cobra.OnInitialize(config.Init)

	rootCmd := root.NewRootCommand()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
