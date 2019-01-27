package main

import (
	"awesomeProject/m/crawler/engine"
	"awesomeProject/m/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
