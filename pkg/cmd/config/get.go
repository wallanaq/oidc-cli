package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewGetCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "get",
		Short: "Retrieves a configuration value",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(viper.GetString(args[0]))
		},
	}

}
