package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGetCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "get [key]",
		Short:   "Gets the value of key from the CLI config file",
		Example: "oidc config get foo",
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viper.GetString(args[0]))
		},
	}

}
