package main

import "fmt"

func main() {
	fmt.Println(square2(2))
	fmt.Println(square2(4))
	fmt.Println(square2(3))
}

func square2(x int) (result int) {
	result = x * x
	defer func() {
		if x == 2 || x == 4 {
			result += x
		}
	}()
	fmt.Print("x ")
	return
}
