package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "hello I'm rpc Server,your msg:   " + request
	fmt.Println("request Success")
	return nil
}

func main() {
	// Create a new router
	fmt.Println("Hello World!")
	// Go语言内置了rpc包，可以直接使用 net/rpc包\

	// 1. 实例化一个server
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	// 2. 注册处理逻辑handler
	// Register the service 服务注册
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 3. 启动server
	for {
		conn, _ := listener.Accept()
		// 使用jsonrpc的序列化协议启动server
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))

	}

}
