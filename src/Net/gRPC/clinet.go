package main

import (
	"Net/gRPC/lws"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:1234", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := lws.NewHelloServiceClient(conn)
	reply, err := client.Hello(context.Background(), &lws.String{Value: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.GetValue())
}
