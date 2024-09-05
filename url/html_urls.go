package url

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	node, err := html.Parse(reader)
	if err != nil {
		return []string{}, fmt.Errorf("error parsing html body: %w", err)
	}

	rawURLs := make([]string, 0)

	nodes := make([]*html.Node, 0)
	nodes = append(nodes, node)
	var n *html.Node
	var c *html.Node
	for len(nodes) > 0 {
		n = nodes[0]
		nodes = nodes[1:]
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				rawURLs = append(rawURLs, attr.Val)
			}
		}
		c = n.FirstChild
		for c != nil {
			nodes = append(nodes, c)
			c = c.NextSibling
		}
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, fmt.Errorf("Error parsing rawBaseURL: %w", err)
	}
	fullURLs := make([]string, 0, len(rawURLs))
	for _, r := range rawURLs {
		u, err := url.Parse(r)
		if err != nil {
			continue // ignore unparsable urls
		}
		if u.IsAbs() {
			fullURLs = append(fullURLs, r)
		} else {
			fullURLs = append(fullURLs, baseURL.ResolveReference(u).String())
		}
	}

	return fullURLs, nil
}
