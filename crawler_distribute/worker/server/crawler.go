package main

import (
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/crawler_distribute/worker"
	"github.com/jormin/go-study/modules/log"
)

func main() {
	log.Fatal(rpcsupport.ServerRpc(config.CrawlHost, &worker.CrawlService{}).Error())
}
