package version

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/mod/semver"
)

var (
	Current  string
	Template string = "oidc-cli version %s (%s/%s)\n"
)

const (
	repoName        = "wallanaq/oidc-cli"
	repoReleasesURL = "https://api.github.com/repos/" + repoName + "/releases/latest"
)

type ReleaseInfo struct {
	TagName     string    `json:"tag_name"`
	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
}

func CheckForUpdate(ctx context.Context) (string, bool, error) {

	info, err := fetchLatestReleaseInfo(ctx)
	if err != nil {
		return "", false, fmt.Errorf("failed to fetch latest release info: %w", err)
	}

	latestVersion := info.TagName

	if semver.Compare(latestVersion, Current) > 0 {
		return latestVersion, true, nil
	}

	return latestVersion, false, nil

}

func fetchLatestReleaseInfo(ctx context.Context) (*ReleaseInfo, error) {

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, repoReleasesURL, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to request latest release info: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	var info ReleaseInfo
	if err := json.NewDecoder(res.Body).Decode(&info); err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	if info.TagName == "" {
		return nil, fmt.Errorf("empty tag name")
	}

	return &info, nil

}
