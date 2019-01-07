package main

import "fmt"

// literal
func main() {
	printMsg("Calling a function")
	showMsg := func(msg string) {
		fmt.Println(msg)
	}
	showMsg("Using a function literal!")
	fmt.Printf("%T\n", showMsg)
	func(msg string) {
		fmt.Println(msg)
	}("quickly reacting")
}

func printMsg(msg string) {
	fmt.Println(msg)
}
