package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

/*
The json package has an Unmarshal() function that
supports decoding data from a byte slice into values.
If the decoded values are assigned to struct fields,
the field names must be exported and therefore have an upper case first letter;
*/

type Inventory struct {
	Total   int      `json:"total"`
	Devices []string `json:"devices"`
}

func main() {
	data, _ := ioutil.ReadFile("file.json")
	inv := Inventory{}
	_ = json.Unmarshal(data, &inv)
	fmt.Printf("Total: %d\nDevices: %q\n", inv.Total, inv.Devices)
}
