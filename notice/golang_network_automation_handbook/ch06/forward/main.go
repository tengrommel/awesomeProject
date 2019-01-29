package main

import (
	"fmt"
	"net"
)

/*
The net.LookupIP() function accept a string and returns a slice of net.IP objects
that represent that host's IPv4 and/or IPv6 address.
*/

func main() {
	ips, _ := net.LookupIP("google.com")
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}
