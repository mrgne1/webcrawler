package url

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeURL(rawUrl string) (string, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return "", fmt.Errorf("Error parsing url: %v", err)
	}

	norm := u.Hostname() + u.Path
	norm, _ = strings.CutSuffix(norm, "/")
	return norm, nil
}
