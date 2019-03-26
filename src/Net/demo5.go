package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/*
*	通过rpc.Dial拨号RPC服务，然后通过client.Call调用具体的RPC方法。
*	在调用client.Call时，
*	第⼀个参数是用点号链接的RPC服务名字和方法名字，
*	第二和第三个参数分别我们定义RPC方法的两个参数。
 */
func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}
