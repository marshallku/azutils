package version

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GitHubRelease struct {
	TagName string `json:"tag_name"`
}

func CheckLatestVersion() (string, error) {
	resp, err := http.Get("https://api.github.com/repos/marshallku/azutils/releases/latest")
	if err != nil {
		return "", fmt.Errorf("failed to fetch latest release: %w", err)
	}
	defer resp.Body.Close()

	var release GitHubRelease
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return TrimVersionPrefix(release.TagName), nil
}
