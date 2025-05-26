package update

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/version"
)

func NewUpdateCheckCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "update-check",
		Short: "Check for updates",
		RunE: func(cmd *cobra.Command, args []string) error {

			cmd.Println("Checking for updates...")

			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			defer cancel()

			latestVersion, updateAvailable, err := version.CheckForUpdate(ctx)
			if err != nil {
				return fmt.Errorf("error checking for updates: %w", err)
			}

			if !updateAvailable {
				cmd.Printf("Up to date! (version %s)\n", version.Current)
				return nil
			}

			cmd.Printf("A new version is available: %s (current: %s)\n", latestVersion, version.Current)

			return nil

		},
	}

}
