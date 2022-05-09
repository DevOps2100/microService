package main

import (
	"fmt"
	"newrpc/client_proxy"
)

func main() {
	// 1. 建立连接
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")

	// var reply = new(string)
	var reply = "123123" // 记住这里一定要服务端匹配的指针类型数据
	err := client.Hello("hello rpc alinx", &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)
}
