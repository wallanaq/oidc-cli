package oidc

import (
	"context"
	"net/http"
)

type OIDCProviderOptions struct {
	Context    context.Context
	HttpClient *http.Client
}

func NewOIDCProvider(cfg *OIDCProviderOptions) OIDCProvider {
	if cfg.HttpClient == nil {
		cfg.HttpClient = &http.Client{}
	}
	if cfg.Context == nil {
		cfg.Context = context.Background()
	}

	return &OIDCProviderImpl{
		ctx:        cfg.Context,
		httpClient: cfg.HttpClient,
	}
}
