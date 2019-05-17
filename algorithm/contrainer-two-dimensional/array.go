package main

import "fmt"

func main() {
	var arr = [4][5] int{
		{4,5,7,8,9},
		{1,2,4,5,6},
		{9,10,11,12,14},
		{3,5,6,8,9},
	}
	var value int = arr[2][3]

	var matrix = [1][3] int{
		{1,2,3},
	}
	fmt.Println("value:", value)
	fmt.Println("matrix:", matrix)
}
