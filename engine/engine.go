package engine

import (
	"log"
	"newcrawler/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	limit := 0
	for len(requests) > 0 {
		if limit > 10 {
			return
		}
		r := requests[0]
		requests = requests[1:]

		log.Printf("Fetching %s", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher:error "+"fetching url %s :%v", r.Url, err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
		limit++
	}
}
