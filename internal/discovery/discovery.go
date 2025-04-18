package discovery

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/wallanaq/oidc-cli/pkg/fileutils"
	"github.com/wallanaq/oidc-cli/pkg/hash"
	"github.com/wallanaq/oidc-cli/pkg/pathutils"
)

// FetchDiscoveryDocument sends an HTTP GET request to the given OIDC discovery URL,
// validates the response, and returns the formatted JSON discovery document.
//
// The response is read, unmarshaled into a generic map for validation,
// and re-encoded as indented JSON for readability.
//
// It logs key steps such as request execution, response status, and parsing success.
//
// Returns an error if the request fails, the response is invalid,
// or the JSON cannot be parsed and formatted.
func FetchDiscoveryDocument(issuer string) ([]byte, error) {

	discoveryURL := strings.TrimRight(issuer, "/") + "/.well-known/openid-configuration"

	log.Printf("Fetching discovery document from: %s", discoveryURL)

	resp, err := http.Get(discoveryURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch discovery document: %w", err)
	}

	log.Printf("Response status: %s", resp.Status)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("OIDC provider returned non-200 status: %s", resp.Status)
	}

	log.Printf("Discovery document fetched successfully")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var pretty map[string]any

	if err := json.Unmarshal(body, &pretty); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	prettyJSON, err := json.MarshalIndent(pretty, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSON: %w", err)
	}

	log.Printf("successfully marshaled JSON")

	return prettyJSON, nil

}

// SaveDiscoveryDocument saves the provided discovery document to a file in the user's
// configuration directory. The filename is derived from the SHA1 hash of the issuer URL,
// ensuring uniqueness.
//
// It logs the saving process, including the destination path.
//
// Returns an error if it fails to get the user's configuration directory or save the file.
func SaveDiscoveryDocument(issuer string, document []byte) error {

	log.Println("Saving discovery document...")

	home, err := pathutils.GetUserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get user config dir for saving: %w", err)
	}

	filename := hash.SHA1Hex(issuer) + ".json"
	path := filepath.Join(home, filename)

	log.Printf("Saving discovery document for issuer '%s' to: %s", issuer, path)

	if err := fileutils.Save(path, document); err != nil {
		return fmt.Errorf("failed to cache discovery document to %s: %w", path, err)
	}

	log.Printf("Discovery document successfully saved.")

	return nil

}
