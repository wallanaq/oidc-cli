package auth

import "github.com/spf13/cobra"

func NewLogoutCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Ends the session (RP-initiated logout)",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
