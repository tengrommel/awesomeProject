package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Sever struct {
	Id int64 `json:"id"`
	Address string `json:"address"`
	Olist []string `json:"olist"`
}

func main() {
	serverList := []Sever{
		{1, "127.0.0.1", []string{"1", "2"}},
		{2, "127.0.0.2", []string{"3", "4"}},
	}
	serverByte, _ := json.Marshal(serverList)
	fmt.Println(string(serverByte))
	fw, err := os.Create("ii.json")
	if err != nil {
		fmt.Println("err:", err)
	}
	fw.Write(serverByte)
	defer fw.Close()
}
