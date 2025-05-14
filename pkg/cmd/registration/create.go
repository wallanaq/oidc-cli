package registration

import "github.com/spf13/cobra"

func NewCreateCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Registers a new client dynamically (RFC 7591)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
