package scheduler

import "newcrawler/engine"

type SimpleScheduler struct {
	requestChan chan engine.Request
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.requestChan <- r }()
}

func (s *SimpleScheduler) ConfigureRequestChan(c chan engine.Request) {
	s.requestChan = c
}
