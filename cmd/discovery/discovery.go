package discovery

import "github.com/spf13/cobra"

var DiscoveryGroup = &cobra.Group{
	ID:    "discovery",
	Title: "Discovery Commands:",
}

var DiscoveryCmd = &cobra.Command{
	Use:     "discovery",
	GroupID: "discovery",
	Short:   "OIDC discovery operations",
	Long:    "Fetch and inspect OpenID Connect discovery documents",
}

func init() {
	DiscoveryCmd.AddCommand(FetchCmd)
}
