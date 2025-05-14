package registration

import (
	"github.com/spf13/cobra"
)

func NewRegistrationCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "registration",
		Short:   "OpenID Connect Dynamic Client Registration",
		GroupID: "oidc",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	return cmd

}
