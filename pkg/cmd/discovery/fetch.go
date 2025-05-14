package discovery

import (
	"github.com/spf13/cobra"
)

func NewFetchCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "fetch",
		Short:   "Fetches metadata from the issuer's .well-known URL",
		Example: "oidc discovery fetch http://example.com/realms/master",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
