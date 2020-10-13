package scheduler

import "github.com/jormin/go-study/crawler/engine"

type QueuedScheduler struct {
	requestCh chan engine.Request
	workerCh  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	s.requestCh <- r
}

func (s *QueuedScheduler) WorkerReady(ch chan engine.Request) {
	s.workerCh <- ch
}

func (s *QueuedScheduler) Run() {
	s.requestCh = make(chan engine.Request)
	s.workerCh = make(chan chan engine.Request)
	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var acticeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				acticeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestCh:
				// send r to a
				requestQ = append(requestQ, r)
			case w := <-s.workerCh:
				// send next request to
				workerQ = append(workerQ, w)
			case acticeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}
