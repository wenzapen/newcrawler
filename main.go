package main

import (
	"newcrawler/engine"
	"newcrawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:       "http://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
	// fetcher.Fetch("https://www.zhenai.com/zhenghun/ankang")

}
