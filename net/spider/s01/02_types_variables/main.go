package main

import "fmt"

func main() {
	var i int
	fmt.Println("i =", i)
	j := 2
	fmt.Printf("%d - %T\n", j, j)
	var k1 uint8 = 20
	fmt.Printf("%d - %T\n", k1, k1)
	var k2 uint16 = 30
	k2 = uint16(k1)
	fmt.Printf("%d - %T\n", k2, k2)
	fmt.Println(uint64(321325 * 424521))
	// compiler error - overflow
}
