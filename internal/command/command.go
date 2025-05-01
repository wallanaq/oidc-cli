package command

import (
	"log"

	"github.com/wallanaq/oidc-cli/internal/oidc"
)

type GlobalFlags struct {
	Verbose bool
	Output  string
}

type Options struct {
	Logger       *log.Logger
	Flags        GlobalFlags
	OIDCProvider oidc.OIDCProvider
}
