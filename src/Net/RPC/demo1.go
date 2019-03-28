package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

/*
*	基于RPC实现watch功能服务端
 */

type KVStoreService struct {
	m      map[string]string           //map类型，用于存储kv数据
	filter map[string]func(key string) //filter成员对应每个Watch调用时定义的过滤器函数列表
	mu     sync.Mutex                  //互斥锁，用于多个goroutine访问或修改时对其它成员提供保护
}

/*实例化对象*/
func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

/*根据key获取value*/
func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}
	return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()
	key, value := kv[0], kv[1] //从入参中获取要设置key、value
	fmt.Println(key, value)
	if oldValue := p.m[key]; oldValue != value { //更新给定key的value
		for _, fn := range p.filter { //获取此key在filter中的fun
			fn(key) //执行fun(key string){ch <- key}，即把要更新值的key写入channel
		}
	}
	p.m[key] = value //更新value
	return nil
}

func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	ch := make(chan string, 10) // buffered
	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()
	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout") //超时等待时间
	case key := <-ch: //从channel中读取更新了value的key
		*keyChanged = key
		return nil
	}
	return nil
}

func main() {
	rpc.Register(NewKVStoreService())
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("ERROR:", err)
			return
		}
		rpc.ServeConn(conn)
	}
}
