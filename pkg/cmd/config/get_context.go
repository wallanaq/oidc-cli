package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGetContextCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "get-context",
		Short:   "Displays the current context",
		Example: "oidc config get-context",
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viper.GetString("context"))
		},
	}

}
