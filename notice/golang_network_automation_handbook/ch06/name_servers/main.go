package main

import (
	"fmt"
	"net"
)

/*
The net.LookupNS() function accepts a domain name as a string
and returns DNS NS records as a slice of NS structs.
*/

func main() {
	nss, _ := net.LookupNS("gmail.com")
	for _, ns := range nss {
		fmt.Println(ns.Host)
	}
}
