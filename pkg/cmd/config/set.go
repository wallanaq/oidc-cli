package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewSetCmd() *cobra.Command {

	return &cobra.Command{
		Use:     "set",
		Short:   "Sets global or context-specific configuration",
		Example: "oidc config set [key] [value]",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.Set(args[0], args[1])
			return viper.WriteConfig()
		},
	}

}
