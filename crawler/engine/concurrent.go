package engine

import (
	"github.com/jormin/go-study/modules/log"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorlerChan(chan Request)
	WorkerReady(chan Request)
	Run()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	count := 0
	for {
		result := <-out

		for _, item := range result.Items {
			log.Info("Got Item #%d: %s", count, item)
			count++
		}

		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}
	}

}

func createWorker(out chan ParseResult, s Scheduler) {
	in := make(chan Request)
	go func() {
		for {
			s.WorkerReady(in)
			r := <-in
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
