package main

import (
	"fmt"
	"net"
)

/*
The net.LookupAddr() function performs a reverse lookup for the address and returns a list map to the address.
*/

func main() {
	names, _ := net.LookupAddr("8.8.8.8")
	for _, name := range names {
		fmt.Println(name)
	}
}
