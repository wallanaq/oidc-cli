package jwt

import "github.com/spf13/cobra"

func NewRevokeCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "revoke",
		Short: "Revokes access or refresh tokens (RFC 7009)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
