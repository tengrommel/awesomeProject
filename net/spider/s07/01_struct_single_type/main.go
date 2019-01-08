package main

import "fmt"

func main() {
	type myType float64
	var total myType

	total = 44
	fmt.Printf("%.2f %T \n", total, total)

	var total2 float64
	fmt.Printf("%.2f %T \n", total2, total2)
}
