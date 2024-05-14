package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureRequestChan(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParseResult)
	c.Scheduler.ConfigureRequestChan(in)
	for i := 0; i < c.WorkCount; i++ {
		go createWorker(in, out)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}
	itemCount := 0
	for {
		result := <-out

		for _, item := range result.Items {
			log.Printf("Got item #%d: %v\n", itemCount, item)
			itemCount++
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult) {
	for {
		request := <-in
		result, err := worker(request)
		if err != nil {
			continue
		}
		out <- result
	}
}
