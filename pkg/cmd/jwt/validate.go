package jwt

import "github.com/spf13/cobra"

func NewValidateCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "validate",
		Short: "Validates JWT signature and claims (JWK - RFC 7517, RFC 7515)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
