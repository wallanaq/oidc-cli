package oidc

import "encoding/json"

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
