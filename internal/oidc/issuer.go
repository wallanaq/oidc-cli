package oidc

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
)

type Issuer struct {
	Value string
}

func (i Issuer) Validate() error {

	if i.Value == "" {
		return errors.New("issuer is required")
	}

	_, err := url.ParseRequestURI(i.Value)
	if err != nil {
		return fmt.Errorf("invalid issuer URL: %w", err)
	}

	// if u.Scheme != "https" {
	// 	return fmt.Errorf("issuer must use https")
	// }

	return nil

}

func (i Issuer) Hash() string {

	sum := sha256.Sum256([]byte(i.Value))

	return hex.EncodeToString(sum[:])

}
