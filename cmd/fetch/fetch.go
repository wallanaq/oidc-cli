package fetch

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wallanaq/oidc-cli/internal/command"
	"github.com/wallanaq/oidc-cli/internal/output"
)

type FetchFlags struct {
	Issuer string
}

func NewFetchCommand(opts *command.Options) *cobra.Command {

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

func run(_ *cobra.Command, _ []string, flags *FetchFlags, opts *command.Options) error {

	provider := opts.OIDCProvider

	oidcConfig, err := provider.FetchConfiguration(flags.Issuer)
	if err != nil {
		return fmt.Errorf("fetch error: %w", err)
	}

	data, err := oidcConfig.MarshalPretty()
	if err != nil {
		return fmt.Errorf("marshal error: %w", err)
	}

	writer, err := output.NewWriter(opts.Flags.Output)
	if err != nil {
		return fmt.Errorf("output writer error: %w", err)
	}

	return writer.Write(data)

}
