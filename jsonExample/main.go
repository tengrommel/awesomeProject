package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	packet := "packt"
	jsonPack, ok := json.Marshal(packet)
	if ok!=nil {
		panic("Could not marshal object")
	}
	fmt.Println(string(jsonPack))
}
