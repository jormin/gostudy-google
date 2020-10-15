package client

import (
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/crawler_distribute/worker"
	"github.com/jormin/go-study/modules/log"
)

func NewClient(host string) (engine.Processor, error) {
	// connect saver service
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}
	return func(r engine.Request) (engine.ParseResult, error) {
		request := worker.SerializeRequest(r)
		res := worker.ParseResult{}
		err = client.Call(config.CrawlRpc, request, &res)
		if err != nil {
			log.Error("Crawler: process error %v", err)
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(res), nil
	}, nil
}
