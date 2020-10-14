package main

import (
	"github.com/jormin/go-study/crawler/zhenai/parser"
	"github.com/jormin/go-study/crawler_distribute/rpcsupport"
	"github.com/jormin/go-study/modules/log"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestRpcServer(t *testing.T) {
	host := ":1234"
	index := "test"
	// start saver service
	go func() {
		err := ServerRpc(host, index)
		if err != nil {
			panic(err)
		}
	}()
	time.Sleep(time.Second)

	// connect saver service
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}

	b, _ := ioutil.ReadFile("./city.html")
	result := parser.ParseUserList(string(b))
	for _, item := range result.Items {
		log.Info("%+v", item)
		log.Info("%+v", reflect.TypeOf(item.Data))
		id := ""
		err = client.Call("SaverService.Save", item, &id)
		if err != nil {
			t.Errorf("error: %v", err)
		}
		log.Info("id: %s", id)
	}

}
