package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wallanaq/oidc-cli/cmd/root"
	"github.com/wallanaq/oidc-cli/internal/cli"
	"github.com/wallanaq/oidc-cli/internal/oidc"
)

func main() {

	opts := &cli.Options{
		Logger: log.New(os.Stdout, "[oidc-cli]", log.LstdFlags),
		OIDCProvider: oidc.NewOIDCProvider(&oidc.ProviderOptions{
			HttpClient: &http.Client{
				Timeout: 5 * time.Second,
			},
		}),
	}

	rootCmd := root.NewRootCommand(opts)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
