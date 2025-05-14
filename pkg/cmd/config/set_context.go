package config

import "github.com/spf13/cobra"

func NewSetContextCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "set-context",
		Short: "Sets the current issuer context",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
