package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

/*
* http协议上的RPC服务端demo
* curl localhost:1234/jsonrpc -X POST --data '{"method":"HelloService.Hello","params":["hello"],"id":0}'
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

//RPC的服务架设在“/jsonrpc”路径，在处理函数中基于http.ResponseWriter和http.Request类型的参数
//构造⼀个io.ReadWriteCloser类型的conn通道。然后基于conn构建针对服务端的json编码解码器。最
//后通过rpc.ServeRequest函数为每次请求处理⼀次RPC方法调用。
func main() {
	rpc.RegisterName("HelloService", new(HelloService))
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234", nil)
}
