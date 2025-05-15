package config

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewSetContextCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "set-context [issuer]",
		Short: "Sets the current issuer context",
		Args:  cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if _, err := url.ParseRequestURI(args[0]); err != nil {
				return fmt.Errorf("invalid issuer URL: %w", err)
			}
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			viper.Set("context", args[0])
			return viper.WriteConfig()
		},
	}

}
