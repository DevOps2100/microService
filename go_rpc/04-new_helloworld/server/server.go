package main

import (
	"fmt"
	"net"
	"net/rpc"
	"newrpc/handler"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello I'm rpc Server,your msg:   " + request
	return nil
}

func main() {
	// Create a new router
	fmt.Println("Hello World!")
	// Go语言内置了rpc包，可以直接使用 net/rpc包\

	// 1. 实例化一个server
	listener, _ := net.Listen("tcp", ":1234")

	// 2. 注册处理逻辑handler
	// Register the service 服务注册
	_ = rpc.RegisterName(handler.HelloServiceName, &HelloService{})
	// 3. 启动server
	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
