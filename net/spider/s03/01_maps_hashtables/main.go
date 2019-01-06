package main

import "fmt"

/*
This example is only to demonstrate how hashtable work
The logic to make hashkeys is not optimized
This solution only works for up to 3 characters.
*/

func main() {
	values := []string{"ABC", "ACB", "BAC", "BCA", "CAB", "CBA"}
	factor := []int{100, 10, 1}
	hashKey := 0
	for v := range values {
		bytes := []byte(values[v])
		f := 0
		hashKey = 0
		for i := range bytes {
			fmt.Print(bytes[i], " ")
			hashKey += int(bytes[i]) * factor[f]
			f++
		}
		fmt.Printf(" (hashKey: %d) \n", hashKey)
	}
}
