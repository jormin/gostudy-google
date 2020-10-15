package persist

import (
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/persist/client"
	"github.com/jormin/go-study/helper"
)

type SimpleSaver struct {
	ItemCh   chan engine.Item
	WorkerCh chan chan engine.Item
}

func (s *SimpleSaver) Run() {
	itemCh, err := client.Saver(config.SaverHost)
	if err != nil {
		helper.LogError("Get saver error", err)
		return
	}
	s.ItemCh = itemCh
}

func (s *SimpleSaver) Submit(item engine.Item) {
	s.ItemCh <- item
}

func (s *SimpleSaver) WorkerChan() chan engine.Item {
	return make(chan engine.Item)
}

func (s *SimpleSaver) WorkerReady(ch chan engine.Item) {
	s.WorkerCh <- ch
}
