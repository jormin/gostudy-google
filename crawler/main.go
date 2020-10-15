package main

import (
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/persist"
	"github.com/jormin/go-study/crawler/scheduler"
	"github.com/jormin/go-study/crawler/zhenai/parser"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/worker/client"
)

func main() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:       "https://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	processor, err := client.NewClient(config.CrawlHost)
	if err != nil {
		panic(err)
	}
	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		Saver:       &persist.SimpleSaver{},
		Processor:   processor,
		WorkerCount: 100,
		Urls:        make(map[string]interface{}),
		Users:       make(map[int]interface{}),
	}
	concurrentEngine.Run(engine.Request{
		Url:    "https://www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})

	//engine.Run(engine.Request{
	//	Url:       "https://album.zhenai.com/u/1079404336",
	//	ParseFunc: parser.ParseProfile,
	//})

	//b, err := ioutil.ReadFile("zhenai/parser/profile.html")
	//if err != nil {
	//	panic(err)
	//}
	//parser.ParseProfile(string(b))

}
