package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan interface{}
}

type Scheduler interface {
	Submit(Request)
	WorkerChan() chan Request
	ReadyNotifier
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	log.Println("engine is running")

	go c.Scheduler.Run()

	out := make(chan ParseResult)

	for i := 0; i < c.WorkCount; i++ {
		go c.createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out

		for _, item := range result.Items {
			// log.Printf("Got item #%d: %v\n", itemCount, item)
			it := item
			go func() {
				c.ItemChan <- it
			}()

		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func (c *ConcurrentEngine) createWorker(workerChan chan Request, out chan ParseResult, ready ReadyNotifier) {
	for {

		ready.WorkerReady(workerChan)
		request := <-workerChan
		result, err := worker(request)
		if err != nil {
			continue
		}
		out <- result
	}
}
