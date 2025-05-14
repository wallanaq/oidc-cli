package main

import (
	"os"

	"github.com/wallanaq/oidc-cli/pkg/cmd/root"
)

func main() {

	rootCmd := root.NewRootCommand()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
