package main

import "fmt"

func powerSeries(a int) (square int, cube int) {
	square =  a*a
	cube = square * a
	return
}

func main() {
	var square int
	var cube int
	square, cube = powerSeries(3)
	fmt.Println("Square ", square, "Cube", cube)
}
