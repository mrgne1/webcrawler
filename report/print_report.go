package report

import (
	"fmt"
	"sort"
)

type pageVisit struct {
	url    string
	visits int
}

func PrintReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %v\n", baseURL)
	fmt.Println("=============================")

	orderedVisits := orderVisits(pages)
	for _, visit := range orderedVisits {
		fmt.Printf("Found %v internal links to %v\n", visit.visits, visit.url)
	}
}

func orderVisits(pages map[string]int) (orderedVisits []pageVisit) {
	orderedVisits = make([]pageVisit, 0, len(pages))

	for url, visits := range pages {
		orderedVisits = append(orderedVisits,
			pageVisit{
				url:    url,
				visits: visits,
			})

	}
	sort.SliceStable(orderedVisits, func(i, j int) bool {
		if orderedVisits[i].visits == orderedVisits[j].visits {
			return orderedVisits[i].url < orderedVisits[j].url
		} else {
			return orderedVisits[i].visits > orderedVisits[j].visits
		}
	})
	return orderedVisits
}
