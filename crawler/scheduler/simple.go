package scheduler

import "github.com/jormin/go-study/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	s.workerChan <- r
}

func (s *SimpleScheduler) ConfigureMasterWorlerChan(ch chan engine.Request) {
	s.workerChan = ch
}
