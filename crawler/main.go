package main

import (
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/scheduler"
	"github.com/jormin/go-study/crawler/zhenai/parser"
)

func main() {

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:       "https://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	concurrentEngine := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
	}
	concurrentEngine.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
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
