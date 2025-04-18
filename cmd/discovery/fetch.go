package discovery

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/command"
	"github.com/wallanaq/oidc-cli/internal/discovery"
)

var (
	save           bool
	issuer, output string
)

var FetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch an OpenID Connect discovery document",
	RunE: func(cmd *cobra.Command, args []string) error {
		return handleFetchCommand(cmd, args)
	},
}

func init() {
	FetchCmd.Flags().StringVarP(&issuer, "issuer", "i", "", "Issuer URL (required)")
	FetchCmd.Flags().BoolVar(&save, "save", false, "Save the discovery document for use by other commands")
	FetchCmd.Flags().StringVar(&output, "output", "", "Output file path or '-' for stdout")

	FetchCmd.MarkFlagRequired("issuer")
}

func handleFetchCommand(_ *cobra.Command, _ []string) error {

	document, err := discovery.FetchDiscoveryDocument(issuer)
	if err != nil {
		return err
	}

	if save {
		err := discovery.SaveDiscoveryDocument(issuer, document)
		if err != nil {
			log.Printf("WARNING: Failed to save/cache discovery document: %v", err)
		}
	}

	return command.HandleOutput(document, output)

}
