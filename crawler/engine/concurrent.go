package engine

import (
	"github.com/jormin/go-study/crawler/model"
	"github.com/jormin/go-study/modules/log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	Saver       Saver
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

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()
	e.Saver.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler, out)
		createSaveWorker(e.Saver)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			if item.Tag == "user" {
				// unique user
				profile := item.Data.(model.SimpleProfile)
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

func createWorker(s Scheduler, out chan ParseResult, ) {
	go func() {
		for {
			in := s.WorkerChan()
			s.WorkerReady(in)
			r := <-in
			result, err := Parse(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func createSaveWorker(s Saver) {
	go func() {
		for {
			in := s.WorkerChan()
			s.WorkerReady(in)
			item := <-in
			id, err := Save(item)
			if err != nil {
				continue
			}
			log.Info("Save item %s", id)
		}
	}()
}
