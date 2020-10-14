package main

import (
	"github.com/jormin/go-study/modules/log"
	rpcdemo "github.com/jormin/go-study/rpc"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	_ = rpc.Register(rpcdemo.DemoRPC{})
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Error("accept error: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}
