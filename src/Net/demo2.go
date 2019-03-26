package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	req, err := http.NewRequest("url", "http://127.0.0.1:3000/time/rfc1123", nil)
	if err != nil {
		log.Println("ERROR", err)
		return
	}
	req.Header.Add("User-Agent", "Gobook Custom User-Agent")
	client := &http.Client{Timeout: time.Second * 3}
	resp, err := client.Do(req)
	if err == nil {
		response, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(response))
		defer resp.Body.Close()
	}
	log.Println("ERROR:", err)

}
