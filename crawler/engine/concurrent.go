package engine

import (
	"encoding/json"
	"fmt"
	"github.com/jormin/go-study/crawler/model"
	"github.com/jormin/go-study/modules/log"
	"io"
	"os"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
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
	// write item to file
	file, err := os.OpenFile("user.txt", os.O_WRONLY|os.O_APPEND, os.ModeAppend)
	defer file.Close()
	if err != nil {
		log.Error("open file error: %v", err)
	}
	// item channel
	itemCh := make(chan Item)
	for i := 0; i < 100; i++ {
		go writerItem(file, itemCh)
	}
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

	count := 0
	for {
		result := <-out
		for _, item := range result.Items {
			if item.Tag == "user" {
				count++
				profile := item.Data.(model.Profile)
				if _, ok := users[profile.BasicInfo.ID]; ok {
					continue
				}
				log.Info("Got Item #%d: %s", count, item)
				users[profile.BasicInfo.ID] = profile
				go func(item Item) {
					itemCh <- item
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

func writerItem(w io.Writer, ch chan Item) {
	for {
		content := <-ch
		b, _ := json.Marshal(content)
		c := fmt.Sprintf("%s\n", b)
		_, err := w.Write([]byte(c))
		if err != nil {
			log.Error("Write item error: %v", err)
		}
	}
}
