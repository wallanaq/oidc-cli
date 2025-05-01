package oidc

import (
	"context"
	"encoding/json"
	"net/http"
)

type OIDCProvider interface {
	FetchConfiguration(issuer string) (*OIDCProviderConfiguration, error)
}

type OIDCProviderOptions struct {
	Context    context.Context
	HttpClient *http.Client
}

type OIDCProviderConfiguration struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserinfoEndpoint      string `json:"userinfo_endpoint"`
	JwksUri               string `json:"jwks_uri,omitempty"`
}

func (c *OIDCProviderConfiguration) MarshalPretty() ([]byte, error) {
	return json.MarshalIndent(c, "", "  ")
}
