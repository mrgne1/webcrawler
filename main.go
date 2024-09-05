package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"webcrawler/report"
	"webcrawler/web"
)

func main() {
	os.Exit(mainCode())
}

func mainCode() int {
	args := os.Args[1:]
	var maxConcurrency int
	var maxPages int
	var err error
	if len(args) < 1 {
		fmt.Println("no website provided")
		return 1
	} else if len(args) == 1 {
		maxConcurrency = 4
		maxPages = 100
	} else if len(args) == 2 { 
		maxConcurrency, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error converting '%v' to int", args[1])
		}
		maxPages = 100
	} else if len(args) == 3 {
		maxConcurrency, err = strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error converting '%v' to int", args[1])
		}
		maxPages, err = strconv.Atoi(args[2])
		if err != nil {
			fmt.Printf("Error converting '%v' to int", args[2])
		}
	} else if len(args) > 3 {
		fmt.Println("too many arguments provided")
		return 1
	}

	crawlURL := args[0]
	fmt.Printf("starting crawl of: %v\n", crawlURL)
	fmt.Printf("maxConcurrency: %v\n", maxConcurrency)
	fmt.Printf("maxPages: %v\n", maxPages)

	baseURL, err := url.Parse(crawlURL)
	if err != nil {
		fmt.Printf("error parsing crawlURL: %s", crawlURL)
		return 1
	}
	var webcrawler web.Config
	webcrawler.Init(baseURL, maxConcurrency, maxPages)

	webcrawler.Wg.Add(1)
	go webcrawler.CrawlPage(crawlURL)
	webcrawler.Wg.Wait()

	report.PrintReport(webcrawler.Pages, crawlURL)

	return 0
}

