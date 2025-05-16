package jwt

import (
	"encoding/json"
)

type JWT struct {
	Raw       string
	Claims    Claims
	Header    map[string]any
	Payload   map[string]any
	signature []byte
}

func (j *JWT) MarshalJSON() ([]byte, error) {
	return json.MarshalIndent(struct {
		Header  map[string]any `json:"header"`
		Payload map[string]any `json:"payload"`
	}{
		Header:  j.Header,
		Payload: j.Payload,
	}, "", "  ")
}
