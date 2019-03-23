package Base

import "fmt"

/*定义结构体类型*/
type Rect struct {
	x, y float64
	width, height float64
}

func (r *Rect) Area() float64 {
	return r.width * r.height
}
/*初始化，实例对象*/
//rect1 := new(Rect)
//rect2 := &Rect{}
//rect3 := &Rect{0, 0, 100, 200}
//rect4 := &Rect{width: 100, height: 200}
/*Go语言中没有构造函数的概念，对象的创建通常交由一个全局的创建函数来完成，以 NewXXX来命名，表示“构造函数”*/
func NewRect(x, y, width, height float64) *Rect {
	return &Rect{x, y, width, height}
}

func main(){
	ret := NewRect(1,2,3,4).Area()
	fmt.Println("ret=",ret)
}