package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wallanaq/oidc-cli/internal/cli"
	"github.com/wallanaq/oidc-cli/internal/oidc"
	"github.com/wallanaq/oidc-cli/pkg/cmd/root"
)

func main() {

	opts := &cli.Options{
		Logger: log.New(os.Stdout, "[oidc-cli]", log.LstdFlags),
		OIDCClient: oidc.NewClient(&oidc.ClientOptions{
			HttpClient: &http.Client{
				Timeout: 3 * time.Second,
			},
		}),
	}

	rootCmd := root.NewRootCommand(opts)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
