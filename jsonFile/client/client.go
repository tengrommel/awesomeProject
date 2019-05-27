package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Sever struct {
	Id int64 `json:"id"`
	Address string `json:"address"`
	Olist []string `json:"olist"`
}

func main() {
	rw,  _ :=os.Open("ii.json")
	stringBytes, _ := ioutil.ReadAll(rw)
	fmt.Println(string(stringBytes))
	servers := []Sever{}
	json.Unmarshal(stringBytes, &servers)
	for i := range servers {
		fmt.Println("Id:", servers[i].Id)
		fmt.Println("Address:", servers[i].Address)
		for j := range servers[i].Olist {
			fmt.Println(servers[i].Olist[j])
		}
	}
}
