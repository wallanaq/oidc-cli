package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewSetCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "set [key] [value]",
		Short:   "Sets an individual value in a CLI config file",
		Example: "oidc config set foo bar",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.Set(args[0], args[1])
			return viper.WriteConfig()
		},
	}

}
