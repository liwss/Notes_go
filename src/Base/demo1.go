package Base

import "fmt"


func main(){
	/*go定义的变量或者导入的包必须使用，否则编译报错*/

	//var v1 int	//定义一个变量

	/*go支持3种定义变量并初始化的方法*/
	var a int = 10	//定义一个变量并初始化
	var b = 10		//定义一个变量并初始化，编译器可以自动推导出类型
	c := 10			//定义一个变量并初始化，编译器可以自动推导出类型
	fmt.Printf("%d %d %d",a,b,c)
	fmt.Printf("\n")

	/*常量定义
	*go语言的常量如果不指定类型，它就如字面意思，是无类型的
	*/
	const d = 2		//无类型浮点常量
	const e int = 5	//整型常量
	const f,g float32 = 3,4 //常量多重赋值，定义float32位的f、g变量，值为3.0和4.0
	fmt.Printf("%f %f",f,g)
	fmt.Printf("\n")

	/*
	*预定义常量true false iota
	*iota:在每一个const关键字出现时被重置为0，然后在下一个const出现之前，没出现一次iota，其值增1
	*const()的写法，类似于C语言的枚举emun
	*/
	const (
		c0 = iota
		c1 = iota
		c2 = iota
		c3
		c4
	)
	fmt.Printf("%d %d %d %d %d",c0,c1,c2,c3,c4)	// 0 1 2 3 4
	fmt.Printf("\n")

	/*
	*数组及切片
	*/
	var arr [10]int = [10]int{0,1,2,3,4,5,6,7,8,9}	//创建一个数组并赋值
	var slice []int = arr[5:]	//基于数组创建一个数组切片
	fmt.Print("数组元素:")
	for _,v := range arr{
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
	fmt.Print("数组切片元素:")
	for _,v := range slice {
		fmt.Print(v, " ")
	}
	fmt.Print("\n")
	//slice1 := make([]int, 5)	//创建一个有5个初始元素的数组切片，元素初始值为0
	slice2 := make([]int, 5, 10) //创建一个有5个初始元素的数组切片，元素初始值为0，并预留10个元素存储空间
	for _, v := range slice2 {
		print(v, " ")
	}
	fmt.Print("\n")
	slice2 = append(slice2,1,2,3,4,5)	//给slice切片后添加元素
	for _, v := range slice2 {
		print(v, " ")
	}

	/*
	* map
	*/
	type personinfo struct {
		ID string
		Name string
		Address string
	}
	var mymap map[string] personinfo	//mymap是声明的map变量名，string是键的类型，personinfo是其中值类型
	//mymap = make(map[string] personinfo)
	//mymap = make(map[string] personinfo,100)	//使用内置函数make创建一个新map，并且指定初始储存能量
	//创建并初始化map
	mymap = map[string] personinfo{
			"1234":personinfo{"1","jack","room 101 ..."},
		}
	fmt.Print(mymap["1234"])

	var mymap1 map[string] string
	mymap1 = map[string] string{
		"1":string("lws"),
		}
	fmt.Print("\n")
	fmt.Print(mymap1["1"])
	mymap1["2"] = string("820")
	value, ok := mymap1["2"]
	if ok{
		fmt.Print("\n")
		fmt.Print(value)
	}
	delete(mymap1,"1")

}
