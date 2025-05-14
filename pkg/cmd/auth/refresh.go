package auth

import "github.com/spf13/cobra"

func NewRefreshCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "refresh",
		Short: "Refreshes tokens using a refresh token",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
