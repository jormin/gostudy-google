package client

import (
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
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
			id := ""
			err := client.Call(config.SaverRpc, item, &id)
			if err != nil {
				continue
			}
			count++
		}
	}()
	return out, nil
}
