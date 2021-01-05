package hello

import (
	"fmt"
	"net/rpc"
)

const ServiceName = "HelloService"

type ServiceClient struct {
	*rpc.Client
}

func (c *ServiceClient) Hello(request string, reply *string) error {
	return c.Client.Call(fmt.Sprintf("%s.Hello", ServiceName), request, reply)
}

func DialHelloService(network, address string) (*ServiceClient, error) {
	client, err := rpc.Dial(network, address)
	if err != nil {
		return nil, err
	}
	return &ServiceClient{client}, nil
}

type ServiceInterface interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(svc ServiceInterface) error {
	return rpc.RegisterName(ServiceName, svc)
}
