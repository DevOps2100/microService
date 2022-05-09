package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	// 1. 建立连接
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}
	// var reply = new(string)
	var reply = "123123" // 记住这里一定要服务端匹配的指针类型数据
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call("HelloService.Hello", "hello rpc alinx", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)

	// 调用
	//  {"method":"HelloService.Hello","params":["hello rpc alinx"],"id":1}

}
