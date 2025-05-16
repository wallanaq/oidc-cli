package jwt

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wallanaq/oidc-cli/internal/jwt"
	"github.com/wallanaq/oidc-cli/internal/output"
)

func NewDecodeCmd() *cobra.Command {

	return &cobra.Command{
		Use:   "decode",
		Short: "Decodes a JWT (access_token or id_token) (JWT - RFC 7519)",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			decoded, err := jwt.Decode(args[0])
			if err != nil {
				return fmt.Errorf("failed to decode jwt: %w", err)
			}

			json, err := decoded.MarshalJSON()
			if err != nil {
				return fmt.Errorf("failed to marshal jwt: %w", err)
			}

			outflag := viper.GetString("output")

			writer := output.NewWriter(outflag)
			writer.Write(json)

			return nil

		},
	}

}
