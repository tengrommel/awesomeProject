package main

import "fmt"

func Identity(order int) [][]float64 {
	var matrix [][]float64
	matrix = make([][]float64, order)
	var i int
	for i=0;i<order;i++ {
		var temp []float64
		temp = make([]float64, order)
		temp[i] = 1
		matrix[i] = temp
	}
	return matrix
}

// add method
func add(matrix1 [2][2]int, matrix2 [2][2]int) [2][2]int {
	var m int
	var l int
	var sum [2][2]int
	for l=0;l<2;l++ {
		for m=0; m<2; m++ {
			sum[l][m] = matrix1[l][m] + matrix2[l][m]
		}
	}
	return sum
}

func main() {
	fmt.Println(Identity(4))
}
