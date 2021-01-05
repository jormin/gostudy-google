package main

import (
	"github.com/jormin/go-study/rpchello2/hello"
	"gitlab.wcxst.com/jormin/go-tools/log"
	"sync"
)

func main() {
	client, err := hello.DialHelloService("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	var requests = []string{
		"rpc", "world", "golang",
	}
	wg := sync.WaitGroup{}
	wg.Add(len(requests))
	for _, request := range requests {
		log.Info("request address: %s, value: %s", &request, request)
		go func(request string) {
			log.Info("request2 address: %s, value: %s", &request, request)
			var reply string
			err = client.Hello(request, &reply)
			if err != nil {
				panic(err)
			}
			log.Info("reply: %s", reply)
			wg.Done()
		}(request)
	}
	wg.Wait()
}
