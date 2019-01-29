package main

import (
	"fmt"
	"net"
)

func main() {
	txts, _ := net.LookupTXT("gmail.com")
	for _, txt := range txts {
		fmt.Println(txt)
	}
}
