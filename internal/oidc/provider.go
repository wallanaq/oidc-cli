package oidc

import (
	"context"
	"net/http"
)

type OIDCProvider interface {
	FetchConfiguration(issuer string) (*ProviderConfiguration, error)
}

type OIDCProviderConfig struct {
	Context    context.Context
	HttpClient *http.Client
}

func NewOIDCProvider(cfg *OIDCProviderConfig) OIDCProvider {
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
