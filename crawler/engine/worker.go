package engine

import (
	"context"
	"encoding/json"
	"github.com/jormin/go-study/crawler/fetcher"
	"github.com/jormin/go-study/crawler/model"
	"github.com/jormin/go-study/modules/log"
	"github.com/olivere/elastic/v7"
	"strconv"
)

func Parse(r Request) (ParseResult, error) {
	log.Info("Fetching %s", r.Url)
	b, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Error("Fetcher: error fetching url %s: %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.Parser.Parse(string(b)), nil
}

func Save(client *elastic.Client, index string, item Item) (id string, err error) {
	profile := item.Data.(model.SimpleProfile)
	b, _ := json.Marshal(item)
	resp, err := client.Index().Index(index).Id(strconv.Itoa(profile.ID)).BodyString(string(b)).Do(context.Background())
	if err != nil {
		log.Error("Index item error: %v", err)
		return id, err
	}
	return resp.Id, nil
}
