package oidc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type OIDCProviderImpl struct {
	ctx        context.Context
	httpClient *http.Client
}

func (p *OIDCProviderImpl) FetchConfiguration(issuer string) (*ProviderConfiguration, error) {

	wellKnown := issuer + "/.well-known/openid-configuration"

	res, err := p.httpClient.Get(wellKnown)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch configuration: %w", err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var config ProviderConfiguration

	if err := json.NewDecoder(res.Body).Decode(&config); err != nil {
		return nil, fmt.Errorf("failed to decode metadata: %w", err)
	}

	return &config, nil

}
