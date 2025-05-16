package jwt

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func Decode(token string) (*JWT, error) {

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid jwt format: expected 3 parts (header, payload, signature)")
	}

	jwt := &JWT{
		Raw: token,
	}

	bHeader, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to decode header: %w", err)
	}

	if err := json.Unmarshal(bHeader, &jwt.Header); err != nil {
		return nil, fmt.Errorf("failed to unmarshal header: %w", err)
	}

	bPayload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	if err := json.Unmarshal(bPayload, &jwt.Payload); err != nil {
		return nil, fmt.Errorf("failed to unmarshal payload: %w", err)
	}

	if err := json.Unmarshal(bPayload, &jwt.StardardClaims); err != nil {
		return nil, fmt.Errorf("failed to unmarshal standard claims: %w", err)
	}

	bSignature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, fmt.Errorf("failed to decode payload: %w", err)
	}

	jwt.signature = bSignature

	return jwt, nil

}
