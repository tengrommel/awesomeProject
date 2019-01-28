package main

import (
	"awesomeProject/m/crawler/engine"
	"awesomeProject/m/crawler/scheduler"
	"awesomeProject/m/crawler/zhenai/parser"
)

func main() {
	e := engine.ConcurrentEngine{Scheduler: &scheduler.SimpleScheduler{}, WorkerCount: 10}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
