package config

import "github.com/spf13/cobra"

func NewSetCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Sets global or context-specific configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
