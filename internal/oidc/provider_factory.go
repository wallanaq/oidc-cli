package oidc

import (
	"context"
	"net/http"
)

type ProviderOptions struct {
	Context    context.Context
	HttpClient *http.Client
}

func NewOIDCProvider(cfg *ProviderOptions) OIDCProvider {
	if cfg.HttpClient == nil {
		cfg.HttpClient = &http.Client{}
	}
	if cfg.Context == nil {
		cfg.Context = context.Background()
	}

	return &oidcProviderImpl{
		ctx:        cfg.Context,
		httpClient: cfg.HttpClient,
	}
}
