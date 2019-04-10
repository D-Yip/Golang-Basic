package main

import (
	"Golang-Basic/crawler/engine"
	"Golang-Basic/crawler/persist"
	"Golang-Basic/crawler/scheduler"
	"Golang-Basic/crawler/zhenai/parser"
)

func main() {

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    persist.ItemSaver(),
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList,
	})
}
