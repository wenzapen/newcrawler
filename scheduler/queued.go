package scheduler

import "newcrawler/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) Submit(r engine.Request) {
	// log.Println("queued scheduler is receiving request")
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Run() {
	s.requestChan = make(chan engine.Request)
	// log.Println("queued scheduler is running")
	s.workerChan = make(chan chan engine.Request)
	requestQ := []engine.Request{}
	workerQ := []chan engine.Request{}
	for {
		// log.Printf("queued scheduler is in for loop: size of reqestQ: %d && sizeof workerQ: %d", len(requestQ), len(workerQ))
		var activeRequest engine.Request
		var activeWorker chan engine.Request
		if len(requestQ) > 0 && len(workerQ) > 0 {
			activeRequest = requestQ[0]
			activeWorker = workerQ[0]
		}
		select {
		case r := <-s.requestChan:
			requestQ = append(requestQ, r)
		case w := <-s.workerChan:
			workerQ = append(workerQ, w)
		case activeWorker <- activeRequest:
			requestQ = requestQ[1:]
			workerQ = workerQ[1:]

		}
	}
}
