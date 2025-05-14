package token

import "github.com/spf13/cobra"

func NewTokenCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "token",
		Short: "Token Utilities",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(
		NewDecodeCmd(),
		NewValidateCmd(),
		NewIntrospectCmd(),
		NewRevokeCmd(),
	)

	return cmd

}
