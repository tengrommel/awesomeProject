package main

import (
	"encoding/json"
	"fmt"
)

type Inventory struct {
	Total   int      `json:"total"`
	Devices []string `json:"devices"`
}

func main() {
	data := &Inventory{
		Total:   3,
		Devices: []string{"SW1", "SW2", "SW3"},
	}
	inv, _ := json.MarshalIndent(data, "", "")
	fmt.Println(string(inv))
}
