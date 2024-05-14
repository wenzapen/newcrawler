package engine

import (
	"log"
	"newcrawler/fetcher"
)

type SimpleEngine struct {
}

func (s *SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	requests = append(requests, seeds...)

	limit := 0
	for len(requests) > 0 {
		if limit > 10000 {
			return
		}
		r := requests[0]
		requests = requests[1:]

		parseResult, err := worker(r)
		if err != nil {
			log.Printf("Fetcher:error "+"fetching url %s :%v", r.Url, err)
			continue
		}

		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}
		limit++
	}
}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
