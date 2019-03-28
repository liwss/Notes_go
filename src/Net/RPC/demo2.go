package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func doClientWork(client *rpc.Client) {
	go func() {
		var keyChanged string //接收m中value有变化的key、
		//先调用watch方法，等待m中value有改变的key，如果有改变，则会通过channel获取，超时时间30s
		err := client.Call("KVStoreService.Watch", 30, &keyChanged)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("watch:", keyChanged)
	}()
	err := client.Call(
		"KVStoreService.Set", [2]string{"abc", "abc-value111111"},
		new(struct{}),
	) //调用Set方法，改变值，如果value变化，会被watch监听到
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 5)
}

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	doClientWork(client)
}
