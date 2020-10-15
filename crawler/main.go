package main

import (
	"flag"
	"fmt"
	"github.com/jormin/go-study/crawler/engine"
	"github.com/jormin/go-study/crawler/persist"
	"github.com/jormin/go-study/crawler/scheduler"
	"github.com/jormin/go-study/crawler/zhenai/parser"
	"github.com/jormin/go-study/crawler_distribute/config"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/crawler_distribute/worker/client"
	"github.com/jormin/go-study/modules/log"
	"net/rpc"
	"strings"
)

var (
	//saverPort    = flag.Int("s", 0, "the port to saver server")
	crawlerPorts = flag.String("c", "", "the ports to crawler server, common separated")
)

func main() {

	flag.Parse()

	//if *saverPort == 0 {
	//	log.Error("Error saver port")
	//	return
	//}

	if *crawlerPorts == "" {
		log.Error("Error crawler ports")
		return
	}

	//engine.SimpleEngine{}.Run(engine.Request{
	//	Url:       "https://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	pool := createProcessorClientPool(strings.Split(*crawlerPorts, ","))

	processor := client.NewClient(pool)

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

// create processor client pool
func createProcessorClientPool(ports []string) chan *rpc.Client {
	var clients []*rpc.Client
	for _, port := range ports {
		c, err := rpcsupport.NewClient(fmt.Sprintf(":%s", port))
		if err != nil {
			log.Error("Processor pool: error connecting to %s: %v", port, err)
			continue
		}
		log.Info("Processor pool: connected to %s", port)
		clients = append(clients, c)
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out
}
