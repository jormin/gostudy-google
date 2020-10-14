package main

import (
	"github.com/jormin/go-study/crawler_distribute/persist"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/modules/log"
	"github.com/olivere/elastic/v7"
)

func main() {
	log.Fatal(ServerRpc(":1234", "profile").Error())
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
