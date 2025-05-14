package token

import "github.com/spf13/cobra"

func NewDecodeCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "decode",
		Short: "Decodes a JWT (access_token or id_token) (JWT - RFC 7519)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
