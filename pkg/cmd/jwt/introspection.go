package jwt

import "github.com/spf13/cobra"

func NewIntrospectCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "introspect",
		Short: "Performs introspection (RFC 7662)",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
