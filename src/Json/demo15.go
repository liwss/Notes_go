package Json

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//从流中读取json
func main() {
	url := "http://127.0.0.1:8000/index"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	defer resp.Body.Close()
	var m map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&m)
	fmt.Println(m)
	fmt.Println(m["name"])
	fmt.Println(m["other"].(map[string]interface{})["phone"])
}