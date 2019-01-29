package main

import (
	"fmt"
	"net"
)

/*
The net.LookupSRV() function accepts a service, protocol, and domain name as a string.
*/

func main() {
	cname, srvs, _ := net.LookupSRV("xmpp-server", "tcp", "google.com")
	fmt.Printf("\ncname: %s \n\n", cname)
	for _, srv := range srvs {
		fmt.Printf("%v:%v:%d:%d\n", srv.Target, srv.Port, srv.Priority, srv.Weight)
	}
}
