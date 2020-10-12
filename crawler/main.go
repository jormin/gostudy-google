package main

import (
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url: "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})

}
