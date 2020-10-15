package engine

import (
	"encoding/json"
	"github.com/jormin/go-study/crawler/model"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	Saver       Saver
	Processor   Processor
	WorkerCount int
	Urls        map[string]interface{}
	Users       map[int]interface{}
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	Run()
	WorkerReady(chan Request)
}

type Saver interface {
	Submit(Item)
	WorkerChan() chan Item
	Run()
	WorkerReady(chan Item)
}

type Processor func(r Request) (ParseResult, error)

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	e.Saver.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.CreateWorker(e.Scheduler, out)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			if item.Tag == "user" {
				// unique user
				b, _ := json.Marshal(item.Data)
				var profile model.SimpleProfile
				_ = json.Unmarshal(b, &profile)
				item.Data = profile
				if _, ok := e.Users[profile.ID]; ok {
					continue
				}
				e.Users[profile.ID] = profile
				e.Saver.Submit(item)
			}
		}

		for _, r := range result.Requests {
			// unique url
			if _, ok := e.Urls[r.Url]; ok {
				continue
			}
			e.Urls[r.Url] = 1
			e.Scheduler.Submit(r)
		}
	}

}

func (e *ConcurrentEngine) CreateWorker(s Scheduler, out chan ParseResult) {
	go func() {
		for {
			in := s.WorkerChan()
			s.WorkerReady(in)
			r := <-in
			result, err := e.Processor(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
