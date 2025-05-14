package config

import "github.com/spf13/cobra"

func NewGetContextCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get-context",
		Short: "Displays the current context",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
