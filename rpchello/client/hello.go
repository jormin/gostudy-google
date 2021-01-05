package main

import (
	"gitlab.wcxst.com/jormin/go-tools/log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "rpc", &reply)
	if err != nil {
		panic(err)
	}
	log.Info("reply: %s", reply)
}
