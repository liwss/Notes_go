package Base

import (
	"errors"
	"fmt"
)
/*
*	函数定义
*
*/
func add(a int, b int) (ret int, err error)  {
	if a<0||b<0{
		err = errors.New("请输入非负数！")
		return
	}
	return a+b,nil
}

func myfunc(args ...interface{}) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

func main()  {
	ret,_ := add(1,2)
	fmt.Print(ret)
	fmt.Print("\n")
	myfunc(1,2,3)
	myfunc(4,"b")
}
