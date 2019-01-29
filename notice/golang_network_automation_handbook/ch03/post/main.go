package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://www.baidu.com"
	data := []byte(`{"hello":"world"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	cookie := http.Cookie{Name: "cookiename", Value: "cookievalue"}
	req.AddCookie(&cookie)
	client := &http.Client{Timeout: time.Second * 10}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)
}
