package oidc

import (
	"context"
	"net/http"
)

type Client interface {
	FetchConfiguration(issuer string) (*ProviderConfiguration, error)
}

type ClientOptions struct {
	Context    context.Context
	HttpClient *http.Client
}

func NewClient(opts *ClientOptions) Client {
	if opts.HttpClient == nil {
		opts.HttpClient = &http.Client{}
	}
	if opts.Context == nil {
		opts.Context = context.Background()
	}

	return &defaultClient{
		ctx:        opts.Context,
		httpClient: opts.HttpClient,
	}
}
