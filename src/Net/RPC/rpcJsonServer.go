package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
*	tcp协议上的跨语言的RPC服务端demo
*	返回给客户端的报文样式  {"id":1,"result":"hello:hello","error":null}
 */

type HelloService struct{}

/*
* Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针
* 类型，并且返回一个error类型，同时必须是公开的方法。
 */
func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	//rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数,所有的注册方法会放在
	//helloService的空间之下
	rpc.RegisterName("HelloService", new(HelloService))
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn)) //传入json编解码器
	}
}
