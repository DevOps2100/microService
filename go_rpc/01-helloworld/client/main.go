package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	// 1. 建立连接
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic("连接失败")
	}
	// var reply = new(string)
	var reply = "123123" // 记住这里一定要服务端匹配的指针类型数据
	err = client.Call("HelloService.Hello", "hello rpc alinx", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
