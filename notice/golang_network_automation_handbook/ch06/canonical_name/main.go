package main

import (
	"fmt"
	"net"
)

/*
The net.LookupCNAME() function accepts a hostname as a string and
returns a single canonical name for the provided host.
*/

func main() {
	cname, _ := net.LookupCNAME("research.swtch.com")
	fmt.Println(cname)
}
