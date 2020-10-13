package engine

import (
	"github.com/jormin/go-study/crawler/fetcher"
	"github.com/jormin/go-study/modules/log"
)

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Info("Got item %s", item)
		}
	}
}

func worker(r Request) (ParseResult, error) {
	log.Info("Fetching %s", r.Url)
	b, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Error("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(string(b)), nil
}
