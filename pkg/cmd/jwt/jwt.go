package jwt

import "github.com/spf13/cobra"

func NewTokenCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "jwt",
		Short: "JWT Utilities",
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
