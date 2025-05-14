package config

import "github.com/spf13/cobra"

func NewConfigCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "config",
		Short: "CLI Configuration & Contexts",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmd.AddCommand(
		NewSetCmd(),
		NewGetCmd(),
		NewSetContextCmd(),
		NewGetContextCmd(),
	)

	return cmd

}
