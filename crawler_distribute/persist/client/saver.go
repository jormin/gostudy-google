package client

import (
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/helper"
	"github.com/jormin/go-study/modules/log"
)

func Saver(host string) (chan engine.Item, error) {
	// connect saver service
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		count := 0
		for {
			item := <-out
			log.Info("Saver: got item %+v", item)
			count++
			id := ""
			err := client.Call(config.SaverRpc, item, &id)
			if err != nil {
				helper.LogError("Saver: save error", err)
			}
			log.Info("Saver: save success: %s", id)
		}
	}()
	return out, nil
}
