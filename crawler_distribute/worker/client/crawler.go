package client

import (
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/worker"
	"github.com/jormin/go-study/modules/log"
	"net/rpc"
)

func NewClient(pool chan *rpc.Client) engine.Processor {
	return func(r engine.Request) (engine.ParseResult, error) {
		request := worker.SerializeRequest(r)
		res := worker.ParseResult{}
		c := <-pool
		err := c.Call(config.CrawlRpc, request, &res)
		if err != nil {
			log.Error("Crawler: process error %v", err)
			return engine.ParseResult{}, err
		}
		return worker.DeserializeResult(res), nil
	}
}
