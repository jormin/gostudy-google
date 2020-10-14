package rpcsupport

import (
	"github.com/jormin/go-study/helper"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func ServerRpc(host string, server interface{}) error {
	err := rpc.Register(server)
	if err != nil {
		return err
	}
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			helper.LogError("accept error", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		helper.LogError("connect error", err)
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
