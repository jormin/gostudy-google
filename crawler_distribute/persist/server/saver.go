package main

import (
	"flag"
	"fmt"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/persist"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/modules/log"
	"github.com/olivere/elastic/v7"
)

var port = flag.Int("port", 0, "the port to start saver server")

func main() {
	flag.Parse()
	if *port == 0 {
		log.Error("must specify a port")
		return
	}
	log.Info("%d", port)
	log.Fatal(ServerRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex).Error())
}

func ServerRpc(host string, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}
	err = rpcsupport.ServerRpc(host, &persist.SaverService{
		ElasticClient: client,
		Index:         index,
	})
	return err
}
