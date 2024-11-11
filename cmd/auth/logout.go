package auth

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/config"
	"github.com/wallanaq/oidc-cli/pkg/database"
)

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Logs out from the Identity Provider, revoking the access token",
	Run: func(cmd *cobra.Command, args []string) {
		handleLogout()
	},
}

func handleLogout() {

	if err := database.Open(); err != nil {
		fmt.Println("Failed to retrieve current authentication.")
		log.Fatalf("Error: %v", err)
	}

	defer database.Close()

	refreshToken, err := database.Get("auth", "refresh_token")

	if err != nil {
		log.Fatalf("Failed to retrieve refresh token from database: %v", err)
	}

	if refreshToken != "" {

		if err := closeUserSession(refreshToken); err != nil {
			fmt.Println("Logout failed.")
			log.Fatalf("Failed to log out using refresh token: %v", err)
		}

		database.Delete("auth", "access_token")
		database.Delete("auth", "refresh_token")

		fmt.Println("Logout successful.")

	}

}

func closeUserSession(token string) error {

	if token == "" {
		return fmt.Errorf("refresh_token is empty")
	}

	data := url.Values{}

	data.Set("refresh_token", token)
	data.Set("client_id", config.ClientID)
	data.Set("client_secret", config.ClientSecret)

	response, err := http.PostForm(config.EndSessionURL, data)

	if err != nil {
		return fmt.Errorf("failed to send logout request: %w", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code received during logout: %d", response.StatusCode)
	}

	return nil

}
