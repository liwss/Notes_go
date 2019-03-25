package main

import (
	"fmt"
	"runtime"
	"sync"
)

//因为runtime.GOMAXPROCS(1)，main函数所在的goroutine占用唯一的线程直到wg.Wait()才挂起。
// 这时循环已经结束i==10，其他routines才有机会依次启动,输出结果都是10
func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}

	wg.Wait()
}
