package main

import "fmt"

func main() {
	slices := [][]int{{2, 3, 4}, {5, 6}, {7, 8, 9, 10}}
	var bigSlice []int
	var p *[][]int
	p = &slices
	for r := range *p {
		bigSlice = append(bigSlice, (*p)[r]...)
	}
	fmt.Printf("%v %T\n", *p, *p)
	fmt.Printf("%v %T\n", bigSlice, bigSlice)
}
