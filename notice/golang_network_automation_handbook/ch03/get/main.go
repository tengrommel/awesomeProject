package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	req.Header.Set("Cache-Control", "no-cache")
	client := &http.Client{Timeout: time.Second * 10}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
}
