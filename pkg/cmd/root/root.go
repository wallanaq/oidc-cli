package root

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wallanaq/oidc-cli/pkg/cmd/auth"
	"github.com/wallanaq/oidc-cli/pkg/cmd/config"
	"github.com/wallanaq/oidc-cli/pkg/cmd/discovery"
	"github.com/wallanaq/oidc-cli/pkg/cmd/registration"
	"github.com/wallanaq/oidc-cli/pkg/cmd/update"
	"github.com/wallanaq/oidc-cli/pkg/cmd/version"
)

func NewRootCommand() *cobra.Command {

	var (
		debug  bool
		output string
	)

	rootCmd := &cobra.Command{
		Use:          "oidc",
		Short:        "OpenID Connect command-line tool",
		Long:         "oidc-cli is a tool for performing Open ID Connect operations",
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if debug {
				slog.SetLogLoggerLevel(slog.LevelDebug)
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	rootCmd.CompletionOptions.DisableDefaultCmd = true

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "enable debug output")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "-", "write file(s) to directory, instead of STDOUT")
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))

	rootCmd.AddGroup(&cobra.Group{
		ID:    "oidc",
		Title: "OpenID Connect Commands:",
	})

	rootCmd.AddCommand(
		auth.NewAuthCommand(),
		config.NewConfigCmd(),
		version.NewVersionCmd(),
		update.NewUpdateCheckCmd(),
		discovery.NewDiscoveryCmd(),
		registration.NewRegistrationCmd(),
	)

	return rootCmd

}
