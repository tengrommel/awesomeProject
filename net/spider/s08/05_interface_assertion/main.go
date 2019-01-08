package main

import "fmt"

func main() {
	var message interface{} = "Hello"
	s, ok := message.(string)
	if ok {
		fmt.Printf("%q %T\n", s, message)
	} else {
		fmt.Printf("value is not a string - %q %T\n", s, message)
	}
	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	f := 20.355
	fmt.Printf("%d %T %T \n", int(f), int(f), f)
	// ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
	var num interface{} = 10
	println(num.(int) + 20)
	fmt.Printf("%d %T \n", num, num)
}
