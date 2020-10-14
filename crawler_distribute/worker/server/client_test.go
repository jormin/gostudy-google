package main

import (
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/crawler_distribute/worker"
	"github.com/jormin/go-study/modules/log"
	"testing"
	"time"
)

func TestCrawlService(t *testing.T) {
	go rpcsupport.ServerRpc(config.CrawlHost, &worker.CrawlService{})
	time.Sleep(time.Second)

	client, err := rpcsupport.NewClient(config.CrawlHost)
	if err != nil {
		panic(err)
	}
	req := worker.Request{
		Url: "http://www.zhenai.com/zhenghun/aba",
		Parser: worker.SerializedParser{
			Name: config.ParseUserList,
			Args: nil,
		},
	}
	res := worker.ParseResult{}
	err = client.Call(config.CrawlRpc, req, &res)
	if err != nil {
		panic(err)
	}
	log.Info("%+v", res.Requests)
	log.Info("%+v", res.Items)
}
