package auth

import "github.com/spf13/cobra"

func NewUserInfoCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "userinfo",
		Short: "User Identity",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
