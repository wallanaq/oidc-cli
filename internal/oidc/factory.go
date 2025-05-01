package oidc

import (
	"context"
	"net/http"
)

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
