package main

import (
	"fmt"
	"net"
)

func main() {
	netGate := "1.1.1.1/27"
	ipv4IP, ipv4Net, _ := net.ParseCIDR(netGate)
	m := ipv4Net.Mask
	dotMask := fmt.Sprintf("%d.%d.%d.%d", m[0], m[1], m[2], m[3])
	fmt.Printf("Dot-decimal notation:%s %s", ipv4IP, dotMask)
	cidrMask := net.IPMask(net.ParseIP(dotMask).To4())
	length, _ := cidrMask.Size()
	slash := fmt.Sprintf("%s/%d", ipv4IP, length)
	fmt.Println("\nCIDR notation:", slash)
}
