package fetch

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/cli"
	"github.com/wallanaq/oidc-cli/internal/issuer"
	"github.com/wallanaq/oidc-cli/internal/output"
)

func NewFetchCommand(opts *cli.Options) *cobra.Command {

	cmd := &cobra.Command{
		Use:     "fetch [issuer]",
		Short:   "Fetch the OpenID Connect configuration",
		Example: "oidc fetch http://example.com/realms/master -o -",
		GroupID: "oidc",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			issuer := &issuer.Issuer{
				Value: args[0],
			}

			if err := issuer.Validate(); err != nil {
				return fmt.Errorf("invalid issuer: %w", err)
			}

			oidcClient := opts.OIDCClient

			oidcConfig, err := oidcClient.FetchConfiguration(issuer.Value)
			if err != nil {
				return fmt.Errorf("fetch error: %w", err)
			}

			data, err := oidcConfig.MarshalPretty()
			if err != nil {
				return fmt.Errorf("marshal error: %w", err)
			}

			writer := output.NewWriter(opts.GlobalFlags.Output)

			return writer.Write(data)

		},
	}

	return cmd

}
