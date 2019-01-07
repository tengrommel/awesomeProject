package main

import "fmt"

func main() {
	addCounter, multCounter := addBy(), multBy()
	fmt.Print(addCounter(2), " ")
	fmt.Print(addCounter(3), " ")
	fmt.Print(addCounter(-1), "\n")
	fmt.Print(multCounter(3), " ")
	fmt.Print(multCounter(4), " ")
	fmt.Print(multCounter(-2))
}

func addBy() func(int) int {
	total := 0
	return func(i int) int {
		total += i
		return total
	}
}

func multBy() func(int) int {
	total := 1
	return func(i int) (ret int) {
		total *= i
		ret = total
		return
	}
}
