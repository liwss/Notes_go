package Base

import (
	"fmt"
	"runtime"
	"sync"
)

var count int = 0
var wg sync.WaitGroup

func main(){
	runtime.GOMAXPROCS(2)
	wg.Add(2)
	go addcount(1)
	go addcount(2)
	wg.Wait()
	fmt.Println(count)
}

func addcount(id int){
	defer wg.Done()
	for i:=0;i<3 ;i++  {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}