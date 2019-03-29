package main

import (
	"net"
	"net/rpc"
	"time"
)

/*
* 反向RPC服务端
* 1：不再提供tcp监听服务，而是主动连接对方服务器
* 2：然后基于tcp链接向对方提供RPC服务
 */

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "hello:" + request
	return nil
}

func main() {
	rpc.Register(new(HelloService))
	for {
		conn, _ := net.Dial("tcp", "localhost:1234")
		if conn == nil {
			time.Sleep(time.Second)
			continue
		}
		rpc.ServeConn(conn)
		conn.Close()
	}
}
