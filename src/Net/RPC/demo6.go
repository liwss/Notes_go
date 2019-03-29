package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/*
* 基于上下文的RPC客户端
 */

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	var ret string
	var reply string
	err = client.Call("HelloService.Login", "password", &ret)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println(reply)
}
