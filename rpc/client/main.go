package main

import (
	"github.com/jormin/go-study/modules/log"
	rpcdemo "github.com/jormin/go-study/rpc"
	"net"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	var result float64
	err = jsonrpc.NewClient(conn).Call("DemoRPC.Div", rpcdemo.Args{
		A: 4,
		B: 0,
	}, &result)
	if err != nil {
		panic(err)
	}
	log.Info("result is %f", result)
}
