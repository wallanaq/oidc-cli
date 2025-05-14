package discovery

import (
	"github.com/spf13/cobra"
)

func NewShowCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "show",
		Short: "Shows saved metadata from the current context",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd

}
