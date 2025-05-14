package discovery

import (
	"github.com/spf13/cobra"
)

func NewDiscoveryCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "discovery",
		Short:   "OpenID Connect Discovery",
		GroupID: "oidc",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(
		NewFetchCmd(),
		NewShowCmd(),
	)

	return cmd

}
