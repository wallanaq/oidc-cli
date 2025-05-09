package oidc

import (
	"context"
	"net/http"
)

type ProviderOptions struct {
	Context    context.Context
	HttpClient *http.Client
}

func NewOIDCProvider(opts *ProviderOptions) OIDCProvider {
	if opts.HttpClient == nil {
		opts.HttpClient = &http.Client{}
	}
	if opts.Context == nil {
		opts.Context = context.Background()
	}

	return &oidcProviderImpl{
		ctx:        opts.Context,
		httpClient: opts.HttpClient,
	}
}
