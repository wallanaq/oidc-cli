package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wallanaq/oidc-cli/cmd/root"
	"github.com/wallanaq/oidc-cli/internal/command"
	"github.com/wallanaq/oidc-cli/internal/oidc"
)

func main() {

	opts := &command.Options{
		Logger: log.New(os.Stdout, "[oidc-cli]", log.LstdFlags),
		OIDCProvider: oidc.NewOIDCProvider(&oidc.OIDCProviderOptions{
			HttpClient: &http.Client{
				Timeout: 5 * time.Second,
			},
		}),
	}

	command := root.NewRootCommand(opts)

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}

}
