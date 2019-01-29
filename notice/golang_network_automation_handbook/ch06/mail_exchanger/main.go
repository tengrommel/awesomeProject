package main

import (
	"fmt"
	"net"
)

/*
The net.LookupMX() function accepts a domain name as string and
returns a slice of MX structs sorted by preference.
*/

func main() {
	mxs, _ := net.LookupMX("google.com")
	for _, mx := range mxs {
		fmt.Println(mx.Host, mx.Pref)
	}
}
