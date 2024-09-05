package url

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name       string
		htmlBody   string
		rawBaseURL string
		expected   []string
	}{
		{
			name:       "empty",
			htmlBody:   "",
			rawBaseURL: "",
			expected:   []string{},
		},
		{
			name: "single path",
			htmlBody: `
<html>
	<body>
		<a href="https://www.google.com">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			rawBaseURL: "",
			expected:   []string{"https://www.google.com"},
		},
		{
			name: "double path",
			htmlBody: `
<html>
	<body>
		<a href="https://www.google.com">
			<span>Boot.dev</span>
		</a>
		<a href="https://www.yahoo.com">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			rawBaseURL: "",
			expected:   []string{"https://www.google.com", "https://www.yahoo.com"},
		},
		{
			name: "relative path",
			htmlBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			rawBaseURL: "https://blog.boot.dev",
			expected:   []string{"https://blog.boot.dev/path/one"},
		},
		{
			name: "relative path and absolute path",
			htmlBody: `
<html>
	<body>
		<a href="/path/one">
			<span>Boot.dev</span>
		</a>
		<a href="http://www.google.com">
			<span>Boot.dev</span>
		</a>
	</body>
</html>
`,
			rawBaseURL: "https://blog.boot.dev",
			expected:   []string{"https://blog.boot.dev/path/one", "http://www.google.com"},
		},
	}

	for i, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(test.htmlBody, test.rawBaseURL)
			if err != nil {
				t.Errorf("Test %v - %v: Unexpected Error: %v", i, test.name, err)
			}
			if !reflect.DeepEqual(actual, test.expected) {
				t.Errorf("Test %v - %v: Error values don't match expected: %v, actual: %v", i, test.name, test.expected, actual)
			}
		})
	}
}
