package auth

import "github.com/spf13/cobra"

func NewLoginCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Starts login flow using Authorization Code + PKCE",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
