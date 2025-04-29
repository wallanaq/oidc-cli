package main

import (
	"os"

	"github.com/wallanaq/oidc-cli/cmd"
)

func main() {
	command := cmd.NewOIDCCommand()
	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
