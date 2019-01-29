package main

import (
	"fmt"
	"net"
)

func main() {
	netGate := "1.1.1.1/30"
	ip, ipNet, _ := net.ParseCIDR(netGate)
	inc := func(ip net.IP) {
		for i := len(ip) - 1; i >= 0; i-- {
			ip[i]++
			if ip[i] >= 1 {
				break
			}
		}
	}
	ipMask := ip.Mask(ipNet.Mask)
	for ip := ipMask; ipNet.Contains(ip); inc(ip) {
		fmt.Println(ip)
	}
}
