package main

import "fmt"

func main() {
	fmt.Printf("%d %o %x %b \n", byte('A'), byte('A'), byte('A'), byte('A'))
	fmt.Printf("%d %[1]o %[1]x %[2]b \n", byte('A'), byte('C'))
	fmt.Printf("%d %o %x %b \n", 'A', 'A', 'A', 'A')
}
