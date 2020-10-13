package engine

import (
	"github.com/jormin/go-study/crawler/fetcher"
	"github.com/jormin/go-study/modules/log"
)

func worker(r Request) (ParseResult, error) {
	log.Info("Fetching %s", r.Url)
	b, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Error("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(string(b)), nil
}
