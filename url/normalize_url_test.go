package url

import (
	"testing"
)

func TestNormalizeURL(t *testing.T) {
	tests := []struct {
		name     string
		inputURL string
		expected string
	}{
		{
			name:     "remove scheme",
			inputURL: "https://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "standard scheme",
			inputURL: "http://blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "trailing slash",
			inputURL: "https://blog.boot.dev/path/",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "query param",
			inputURL: "https://blog.boot.dev/path?q=hellothere",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "username",
			inputURL: "https://asimpleuser@blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "username and password",
			inputURL: "https://asimpleuser:1234@blog.boot.dev/path",
			expected: "blog.boot.dev/path",
		},
		{
			name:     "port",
			inputURL: "https://blog.boot.dev:8080/path",
			expected: "blog.boot.dev/path",
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := normalizeURL(tc.inputURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if actual != tc.expected {
				t.Errorf("Test %v - '%s' FAIL: unexpected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
