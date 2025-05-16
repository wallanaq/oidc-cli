package root

import (
	"log/slog"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/pkg/cmd/auth"
	"github.com/wallanaq/oidc-cli/pkg/cmd/config"
	"github.com/wallanaq/oidc-cli/pkg/cmd/discovery"
	"github.com/wallanaq/oidc-cli/pkg/cmd/jwt"
	"github.com/wallanaq/oidc-cli/pkg/cmd/registration"
)

var (
	version         string = "v0.0.0"
	versionTemplate string = "oidc-cli version {{.Version}}\n"
)

func NewRootCommand() *cobra.Command {

	var debug bool

	rootCmd := &cobra.Command{
		Use:          "oidc",
		Short:        "OpenID Connect command-line tool",
		Long:         "oidc-cli is a tool for performing Open ID Connect operations",
		Version:      version,
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

	rootCmd.SetVersionTemplate(versionTemplate)

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")

	rootCmd.AddGroup(&cobra.Group{
		ID:    "oidc",
		Title: "OpenID Connect Commands:",
	})

	rootCmd.AddCommand(
		discovery.NewDiscoveryCmd(),
		registration.NewRegistrationCmd(),
		auth.NewAuthCommand(),
		jwt.NewTokenCmd(),
		config.NewConfigCmd(),
	)

	return rootCmd

}
