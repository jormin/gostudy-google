package main

import (
	"flag"
	"fmt"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/crawler_distribute/worker"
	"github.com/jormin/go-study/modules/log"
)

var port = flag.Int("port", 0, "the port to start crawler server")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Error("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServerRpc(fmt.Sprintf(":%d", *port), &worker.CrawlService{}).Error())
}
