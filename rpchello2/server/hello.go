package main

import (
	"fmt"
	"github.com/jormin/go-study/rpchello2/hello"
	"gitlab.wcxst.com/jormin/go-tools/log"
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
	hello.RegisterHelloService(new(HelloService))
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("listen tcp error: %v", err)
	}
	for {
		con, err := listener.Accept()
		if err != nil {
			log.Fatal("accept connection error: %v", err)
		}
		go rpc.ServeConn(con)
	}
}
