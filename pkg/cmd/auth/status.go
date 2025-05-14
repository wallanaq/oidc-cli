package auth

import "github.com/spf13/cobra"

func NewStatusCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Displays stored tokens and their status",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
