package web

import (
	"fmt"
	"net/url"
	"sync"
	wcURL "webcrawler/url"
)

type Config struct {
	Pages              map[string]int
	BaseURL            *url.URL
	Mu                 *sync.Mutex
	ConcurrencyControl chan struct{}
	Wg                 *sync.WaitGroup
	MaxPages           int
}

func (cfg *Config) Init(baseURL *url.URL, n int, maxPages int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cfg.BaseURL = baseURL
	cfg.Pages = make(map[string]int)
	cfg.ConcurrencyControl = make(chan struct{}, n)
	cfg.Wg = &wg
	cfg.Mu = &mu
	cfg.MaxPages = maxPages
}

func (cfg *Config) CrawlPage(rawCurrentURL string) {
	cfg.ConcurrencyControl <- struct{}{}
	defer cfg.Wg.Done()
	defer func() { <-cfg.ConcurrencyControl }()
	fmt.Printf("Crawling: %v\n", rawCurrentURL)

	if cfg.MaxPagesReached() {
		return
	}

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		return
	}

	if cfg.BaseURL.Hostname() != currentURL.Hostname() {
		return
	}

	normCurrentURL, err := wcURL.NormalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	isFirst := cfg.AddPageVisit(normCurrentURL)

	if isFirst {
		html, err := GetHTML(rawCurrentURL)
		if err != nil {
			return
		}

		urls, err := wcURL.GetURLsFromHTML(html, cfg.BaseURL.String())
		if err != nil {
			return
		}

		for _, u := range urls {
			cfg.Wg.Add(1)
			go cfg.CrawlPage(u)
		}
	}
}

func (cfg *Config) AddPageVisit(normalizedURL string) (isFirst bool) {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()
	_, ok := cfg.Pages[normalizedURL]
	if ok {
		cfg.Pages[normalizedURL] += 1
	} else {
		cfg.Pages[normalizedURL] = 1
	}
	return !ok
}

func (cfg *Config) MaxPagesReached() bool {
	cfg.Mu.Lock()
	defer cfg.Mu.Unlock()

	return len(cfg.Pages) >= cfg.MaxPages
}
