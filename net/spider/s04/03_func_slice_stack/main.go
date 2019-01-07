package main

import "fmt"

func main() {
	s := []int{}
	s = push(s, 1, 2, 3)
	fmt.Println(s)
	s = pop(s)
	fmt.Println(s)
	s = push(s, 10, 11)
	fmt.Println(s)
	fmt.Println(top(s))
}

func push(s []int, newS ...int) []int {
	return append(s, newS...)
}

func pop(s []int) []int {
	return s[:len(s)-1]
}

func top(s []int) int {
	return s[len(s)-1]
}
