package client_proxy

import (
	"net/rpc"
	"newrpc/handler"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protol, address string) *HelloServiceStub {
	conn, err := rpc.Dial(protol, address)
	if err != nil {
		panic("connect error~")
	}
	return &HelloServiceStub{Client: conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+"hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
