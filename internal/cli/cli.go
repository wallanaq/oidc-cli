package cli

import (
	"log"

	"github.com/wallanaq/oidc-cli/internal/oidc"
)

type GlobalFlags struct {
	Verbose bool
	Output  string
}

type Options struct {
	Flags        GlobalFlags
	Logger       *log.Logger
	OIDCProvider oidc.OIDCProvider
}
