package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Sever struct {
	CorpOrderId string `json:"corpOrderId"`
	AccidentId string `json:"accidentId"`
	AccidentSolutionType string `json:"accidentSolutionType"`
	AmountInCent int `json:"amountInCent"`
	RefundDishCount int `json:"refundDishCount"`
	RefundToCompany int `json:"refundToCompany"`
	RefundToUser int `json:"refundToUser"`
}

func main() {
	rw,  _ :=os.Open("ii.json")
	stringBytes, _ := ioutil.ReadAll(rw)
	fmt.Println(string(stringBytes))
	servers := []Sever{}
	json.Unmarshal(stringBytes, &servers)
	for i := range servers {
		fmt.Println(servers[i])
	}
}
