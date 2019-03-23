package Json

import (
	"encoding/json"
	"fmt"
	"log"
)

//json转换为map
func main() {
	var JSON = `{"age":"20","name":"lws","other":{"addr":"china","phone":"123456"}}`
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)	//json转化为map
	if err != nil {
		log.Println("ERROR:", err)
		return
	}

	fmt.Println("Name:", c["name"])
	fmt.Println("phone:", (c["other"]).(map[string]interface{})["phone"])

}