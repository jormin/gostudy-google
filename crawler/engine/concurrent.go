package engine

import (
	"github.com/jormin/go-study/crawler/model"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemCh      chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	// user  map
	users := make(map[int]interface{})

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			if item.Tag == "user" {
				profile := item.Data.(model.Profile)
				if _, ok := users[profile.BasicInfo.ID]; ok {
					continue
				}
				users[profile.BasicInfo.ID] = profile
				go func(item Item) {
					e.ItemCh <- item
				}(item)
			}
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}

}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			r := <-in
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
