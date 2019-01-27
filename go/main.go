package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

var tr = &http.Transport{
	TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
	DisableCompression: true,
}

var c = &http.Client{
	Timeout:   10 * time.Second,
	Transport: tr,
}

func main() {
	requestStruct := struct {
		Bucket     string `json:"bucket"`
		Key        string `json:"key"`
		ExpireTime string `json:"expire_time"`
	}{
		Bucket:     "f",
		Key:        "f",
		ExpireTime: "123",
	}

	buf := new(bytes.Buffer)
	requestJson, _ := json.Marshal(requestStruct)
	buf.Write(requestJson)
	req, _ := http.NewRequest("POST", "https://g4amuqwmrh.execute-api.cn-north-1.amazonaws.com.cn/prod/uploadurl", buf)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "YrJNbWkyJA4XPJPEmqGm5A9gsDJakiG8rrCJF5lb")
	//re, _ := ioutil.ReadAll(req.Body)
	//fmt.Println(re)
	response, _ := c.Do(req)
	defer response.Body.Close()
	ee, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(ee))

	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(requestStruct)
	fmt.Println(b.Bytes())
	req, _ = http.NewRequest("POST", "https://g4amuqwmrh.execute-api.cn-north-1.amazonaws.com.cn/prod/uploadurl", strings.NewReader(string(b.Bytes())))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", "YrJNbWkyJA4XPJPEmqGm5A9gsDJakiG8rrCJF5lb")
	res, _ := c.Do(req)
	io.Copy(os.Stdout, res.Body)
}
