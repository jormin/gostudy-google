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
	//s.ItemCh = make(chan engine.Item)
	//s.WorkerCh = make(chan chan engine.Item)
	//go func() {
	//	var itemQ []engine.Item
	//	var workerQ []chan engine.Item
	//	for {
	//		var activeItem engine.Item
	//		var activeWorker chan engine.Item
	//		if len(itemQ) > 0 && len(workerQ) > 0 {
	//			activeItem = itemQ[0]
	//			activeWorker = workerQ[0]
	//		}
	//		select {
	//		case item := <-s.ItemCh:
	//			itemQ = append(itemQ, item)
	//		case w := <-s.WorkerCh:
	//			workerQ = append(workerQ, w)
	//		case activeWorker <- activeItem:
	//			itemQ = itemQ[1:]
	//			workerQ = workerQ[1:]
	//		}
	//	}
	//}()
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
