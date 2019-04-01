package main

import (
	"Net/gRPCStream/lws"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

/*
* gRPC server 双向流demo
* gRPC的context.Context参数，为每个方法调用提供上下文支持
 */

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *lws.String) (*lws.String, error) {
	reply := &lws.String{Value: "hello:" + args.GetValue()}
	return reply, nil
}

func (p *HelloServiceImpl) Channel(stream lws.HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF { //表示客户端流被关闭
				return nil
			}
			return err
		}
		reply := &lws.String{Value: "hello:" + args.GetValue()}
		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func main() {
	grpcServer := grpc.NewServer()
	lws.RegisterHelloServiceServer(grpcServer, new(HelloServiceImpl))
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer.Serve(lis)
}
