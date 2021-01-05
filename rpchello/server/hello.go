package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = fmt.Sprintf("hello %s", request)
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	con, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	rpc.ServeConn(con)
}
