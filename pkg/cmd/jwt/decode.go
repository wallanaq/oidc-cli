package jwt

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/jwt"
)

func NewDecodeCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "decode",
		Short: "Decodes a JWT (access_token or id_token) (JWT - RFC 7519)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			jwt, err := jwt.Decode(args[0])
			if err != nil {
				return fmt.Errorf("failed to decode jwt: %w", err)
			}

			json, err := jwt.MarshalJSON()
			if err != nil {
				return fmt.Errorf("failed to marshal jwt: %w", err)
			}

			fmt.Println(string(json))

			return nil

		},
	}

}
