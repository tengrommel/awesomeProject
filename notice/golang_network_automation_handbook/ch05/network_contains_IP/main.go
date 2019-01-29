package main

import (
	"fmt"
	"net"
)

/*
CIDR
*/

func main() {
	netGate := "1.1.1.0/24"
	ips := []string{"1.1.1.1", "2.2.2.2"}
	for _, ip := range ips {
		_, cidrNet, _ := net.ParseCIDR(netGate)
		addr := net.ParseIP(ip)
		result := cidrNet.Contains(addr)
		fmt.Printf("%s contains %s: %t\n", cidrNet, addr, result)
	}
}
