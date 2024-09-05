package web

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("unknown error getting webpage: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 500 {
		return "", fmt.Errorf("server error getting webpage: %v", err)
	} else if res.StatusCode >= 400 {
		return "", fmt.Errorf("client error getting webpage: %v", err)
	}

	content := res.Header.Get("content-type")
	if !strings.Contains(content, "text/html") {
		return "", errors.New(fmt.Sprintf("content-type is '%v' not text/html", content))
	}

	htmlData, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reading html from response: %v", err)
	}

	return string(htmlData), nil
}
