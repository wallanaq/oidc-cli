package fetch

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/cli"
	"github.com/wallanaq/oidc-cli/internal/cli/output"
)

type FetchFlags struct {
	Issuer string
}

func NewFetchCommand(opts *cli.Options) *cobra.Command {

	flags := &FetchFlags{}

	cmd := &cobra.Command{
		Use:     "fetch",
		Short:   "Fetch the OpenID Connect configuration",
		GroupID: "oidc",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd, args, flags, opts)
		},
	}

	flags.Bind(cmd)

	return cmd

}

func (f *FetchFlags) Bind(cmd *cobra.Command) {

	flags := cmd.Flags()

	flags.StringVarP(&f.Issuer, "issuer", "i", "", "Issuer")

	cmd.MarkFlagRequired("issuer")

}

func run(_ *cobra.Command, _ []string, flags *FetchFlags, opts *cli.Options) error {

	provider := opts.OIDCProvider

	meta, err := provider.FetchConfiguration(flags.Issuer)
	if err != nil {
		return fmt.Errorf("fetch error: %w", err)
	}

	data, err := json.MarshalIndent(meta, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	writer, err := output.NewOutputWriter(opts.Flags.Output)
	if err != nil {
		return fmt.Errorf("output writer error: %w", err)
	}

	return writer.Write(data)

}
