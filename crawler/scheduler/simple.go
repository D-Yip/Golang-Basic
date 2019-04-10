package scheduler

import "Golang-Basic/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkChan(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit(r engine.Request) {

	//	send Request down to worker chan
	go func() {
		s.workerChan <- r
	}()
}
