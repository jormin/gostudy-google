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
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	in := make(chan Request)
	out := make(chan ParseResult)

	e.Scheduler.ConfigureMasterWorlerChan(in)

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(in, out)
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

func createWorker(in chan Request, out chan ParseResult) {
	go func() {
		for {
			r := <-in
			result, err := worker(r)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
