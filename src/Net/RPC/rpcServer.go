package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

type Calculator struct{} /*对象类型声明，用于RPC注册*/
/*请求参数类型*/
type CalculatorRequest struct {
	A int
	B int
}

/*响应参数类型*/
type CalculatorResponse struct {
	Result int
}

/*RPC注册对象的方法，参数只能有两个序列化对象，其中第2个参数必须是指针类型*/
func (cal *Calculator) Add(req CalculatorRequest, rsp *CalculatorResponse) error {
	rsp.Result = req.A + req.B
	return nil
}

func (cal *Calculator) Sub(req CalculatorRequest, rsp *CalculatorResponse) error {
	rsp.Result = req.A - req.B
	return nil
}

func main() {
	rpc.Register(new(Calculator))                      //把Calculator对象的所有满足RPC规则的方法注册为RPC函数
	listen, err := net.Listen("tcp", "127.0.0.1:8000") //建立一个tcp连接
	if err != nil {
		log.Println("ERROR:", err)
	}
	fmt.Println("RPC Server start ...")
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("ERROR:", err)
		}
		go rpc.ServeConn(conn) //在tcp连接的基础上为客户端提供RPC服务
	}
}
