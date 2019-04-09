package main

import (
	"Golang-Basic/crawler/engine"
	"Golang-Basic/crawler/zhenai/parser"
)

func main() {

	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList,
	})
}
