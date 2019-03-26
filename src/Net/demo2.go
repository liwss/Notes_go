package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	req, err := http.NewRequest("url", "http://127.0.0.1:3000/time/rfc1123", nil)
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	req.Header.Add("User-Agent", "Gobook Custom User-Agent")
	client := &http.Client{}
	resp, err := client.Do(req)
	response, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(response))
}
