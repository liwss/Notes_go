package Base

import (
	"fmt"
	"runtime"
	"sync"
)

func main(){
	runtime.GOMAXPROCS(2)	//分配一个逻辑处理器给调度器使用
	//runtime.GOMAXPROCS(runtime.NumCPU())  //runtime.NumCPU()获取系统cpu个数
	var wg sync.WaitGroup	//wg表示用来等待程序完成，计数为2，表示等待2个goroutine完成
	wg.Add(2)

	fmt.Println("Start goroutine ......")
	/*打印小写字母表*/
	go func() {
		defer wg.Done()		//defer关键字，函数执行完成退出前调用，wg.done()告诉wg.wait() goroutine执行完成
		for count:=0;count<3 ;count++  {
			for ch:='a';ch<'a'+26 ;ch++  {
				fmt.Printf("%c",ch)
			}
		}
	}()
	/*打印大写字母表*/
	go func() {
			defer wg.Done()		//defer关键字，函数执行完成退出前调用
			for count:=0;count<3 ;count++  {
				for ch:='A';ch<'A'+26 ;ch++  {
					fmt.Printf("%c",ch)
				}
			}
	}()

	fmt.Println("Waiting finish ......")
	wg.Wait()	//等待goroutine运行结束
	fmt.Println("exit ......")
}
