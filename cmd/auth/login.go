package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/config"
	"github.com/wallanaq/oidc-cli/pkg/browser"
	"github.com/wallanaq/oidc-cli/pkg/database"
	"github.com/wallanaq/oidc-cli/pkg/server"
	"golang.org/x/oauth2"
)

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Authenticate and obtain access credentials",
	Run: func(cmd *cobra.Command, args []string) {
		handleLogin()
	},
}

func handleLogin() {

	oauth2Config := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
		Scopes:       []string{},
		Endpoint: oauth2.Endpoint{
			AuthURL:  config.AuthURL,
			TokenURL: config.TokenURL,
		},
	}

	state := "random-state"

	authURL := oauth2Config.AuthCodeURL(state)

	stopChan := make(chan os.Signal, 1)

	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/callback", func(writer http.ResponseWriter, request *http.Request) {

		code := request.URL.Query().Get("code")

		if code == "" {
			http.Error(writer, "Missing code", http.StatusBadRequest)
			return
		}

		token, err := oauth2Config.Exchange(context.Background(), code)

		if err != nil {
			http.Error(writer, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(writer, "Login successful! You can close your browser now.")

		fmt.Println("Login successful!")

		if err := database.Open(); err != nil {
			log.Fatalf("The CLI authentication process could not be completed: %v", err)
			os.Exit(1)
		}

		defer database.Close()

		if err := database.Set("auth", "access_token", token.AccessToken); err != nil {
			log.Fatalf("Error saving access_token: %v", err)
			os.Exit(1)
		}

		if err := database.Set("auth", "refresh_token", token.RefreshToken); err != nil {
			log.Fatalf("Error saving refresh_token: %v", err)
			os.Exit(1)
		}

		close(stopChan)

	})

	go func() {
		if err := server.Start(":8088", nil); err != nil {
			log.Printf("Failed to create callback for authentication")
			log.Fatalf("%v", err)
		}
	}()

	if err := browser.Open(authURL); err != nil {
		fmt.Println("Failed to open the browser for authentication.")
		log.Fatalf("%v", err)
	}

	<-stopChan

	if err := server.Shutdown(); err != nil {
		log.Println("Error during server shutdown:", err)
	}

}
