package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

/*
* 基于上下文的RPC服务端，为客户端每个链接提供独立的RPC服务
 */

type HelloService struct {
	conn    net.Conn //用于识别不同的客户端链接
	isLogin bool     //用于判断是否登录
}

func (p *HelloService) Login(request string, reply *string) error {
	if request != "password" {
		return fmt.Errorf("auth failed")
	}
	log.Println("login ok")
	p.isLogin = true
	return nil
}
func (p *HelloService) Hello(request string, reply *string) error {
	if !p.isLogin {
		return fmt.Errorf("please login")
	}
	*reply = "hello:" + request + ", from" + p.conn.RemoteAddr().String()
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go func() {
			defer conn.Close()
			p := rpc.NewServer()                  //新建一个RPC服务
			p.Register(&HelloService{conn: conn}) //根据不同的客户端链接注册HelloService的RPC服务
			p.ServeConn(conn)                     //启动RPC服务
		}()
	}
}
