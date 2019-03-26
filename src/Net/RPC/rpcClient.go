package main

import (
	"fmt"
	"log"
	"net/rpc"
)

/*请求RPC服务端的结构类型*/
type CalculatorRequest struct {
	A int
	B int
}

/*获取响应的结构类型*/
type CalculatorResponse struct {
	Result int
}

func main() {
	//拨号RPC服务端
	client, err := rpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Println("ERROR:", err)
	}
	req := CalculatorRequest{3, 5}
	var rsp CalculatorResponse
	//调用RPC服务端具体方法，第一个参数为服务端注册对象和调用方法组合
	err = client.Call("Calculator.Sub", req, &rsp)
	if err != nil {
		log.Println("ERROR:", err)
	}
	fmt.Println(req.A, "+", req.B, "=", rsp.Result)
}
