package Base

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main(){
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	num := make(chan int)	//创建一个无缓冲通道
	go show(num,"A")
	go show(num,"B")
	num <- 1
	wg.Wait()
}

func show(num chan int,id string){
	defer wg.Done()
	for{
		ret, ok := <-num
		if !ok{
			fmt.Println("game over!")
			return
		}
		if ret>20{
			close(num)
			return
		}
		fmt.Println(id," ret=",ret)
		ret++
		num <- ret
	}
}