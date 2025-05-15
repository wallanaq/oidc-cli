package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Init() {

	userConfigDir, err := os.UserConfigDir()
	cobra.CheckErr(err)

	cliConfigDir := filepath.Join(userConfigDir, "oidc-cli")

	if err := os.MkdirAll(cliConfigDir, os.ModePerm&0755); err != nil {
		cobra.CheckErr(err)
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cliConfigDir)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SafeWriteConfig()
		}
	}

}
