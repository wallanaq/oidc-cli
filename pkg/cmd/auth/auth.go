package auth

import (
	"github.com/spf13/cobra"
)

func NewAuthCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "auth",
		Short:   "OpenID Connect Core",
		GroupID: "oidc",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(
		NewLoginCmd(),
		NewLogoutCmd(),
		NewRefreshCmd(),
		NewStatusCmd(),
		NewUserInfoCmd(),
	)

	return cmd

}
