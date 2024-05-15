package main

import (
	"newcrawler/engine"
	"newcrawler/persist"
	"newcrawler/scheduler"
	"newcrawler/zhenai/parser"
)

func main() {
	e := &engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkCount: 100,
		ItemChan:  persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
	// fetcher.Fetch("https://www.zhenai.com/zhenghun/ankang")

}
