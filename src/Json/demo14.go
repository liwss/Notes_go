package Json

import (
	"encoding/json"
	"fmt"
	"log"
)

//map转换为json
func main()  {
	mymap := make(map[string]interface{})
	mymap["name"] = "lws"
	mymap["age"] = "20"
	mymap["other"] = map[string]interface{}{
		"addr": "china",
		"phone": "123456",
	}

	//data, err := json.MarshalIndent(mymap, "","  ")  //生成带有缩进格式的json
	data, err := json.Marshal(mymap)				   //生成不带缩进的json
		if err != nil{
		log.Println("ERROR", err)
		return
	}
	fmt.Println(string(data))
}