package config

import "github.com/spf13/cobra"

func NewGetCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Retrieves a configuration value",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
