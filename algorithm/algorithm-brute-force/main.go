package main

import "fmt"

// findElement method given array and k element
func findElement(arr[10] int, k int) bool {
	for i:=0; i<len(arr); i++ {
		if arr[i] == k {
			return true
		}
	}
	fmt.Println(len(arr))
	return false
}

func main() {
	var arr = [10]int{1, 4, 7, 8, 3, 9, 2, 4, 1, 8}
	var check bool = findElement(arr, 10)
	fmt.Println(check)
}
