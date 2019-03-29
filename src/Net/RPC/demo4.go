package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

/*
* 反向RPC客户端
* 1:RPC客户端不再主动去链接服务端，而是在公共地址提供一个tcp服务
* 2：利用服务端链接过来的信息通过channel传送给doClientWork发起RPC调用
 */

func doClientWork(clientChan <-chan *rpc.Client) {
	client := <-clientChan
	defer client.Close()
	var reply string
	err := client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	clientChan := make(chan *rpc.Client)
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Fatal("Accept error:", err)
			}
			clientChan <- rpc.NewClient(conn)
		}
	}()
	doClientWork(clientChan)
}
