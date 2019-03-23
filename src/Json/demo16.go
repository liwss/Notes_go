package Json

import (
	"encoding/json"
	"fmt"
	"log"
)

func main()  {
	type Book struct {
		Title string
		Authors string
		Publisher string
		IsPublished bool
		Price float32
	}

	a := Book{
		"Go语言编程",
		"qqqq",
		"ituring.com.cn",
		true,
		9.99,
	}
	fmt.Println(a)
	b, err := json.Marshal(a)
	if err != nil{
		log.Println("ERROR:",err)
		return
	}
	fmt.Println(string(b))
}
