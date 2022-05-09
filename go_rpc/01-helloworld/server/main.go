package main

import (
	"fmt"
	"net"
	"net/rpc"
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
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	// 2. 注册处理逻辑handler
	// Register the service 服务注册
	_ = rpc.RegisterName("HelloService", &HelloService{})

	// 3. 启动server
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("accept error:", err)
		return
	}
	rpc.ServeConn(conn)

	//  一连串的代码大部分都是net的包，好像和rpc没有关系
	//  可以看看net包的文档

	/*
		rpc调用中有几个问题需要解决
			1.call id
			2.序列化和反序列化

		3.go语言的rpc的序列化和反序列化协议是什么(Gob)
		4.是否能替换常见的序列化
	*/
}
