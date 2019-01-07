package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(len("Good Day1"))
	fmt.Println(reflect.TypeOf("你好 Day"[0]))
	fmt.Println("你好 Day")
	fmt.Printf("%v\n", "你好 Day"[0])
	fmt.Println()

	s := "SuperHelpful"
	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:9])
	fmt.Println(s[:])

	fmt.Println(s[:5] + s[5:])
	fmt.Println(&s)
	// s[0] = 'D' // compile error:
	s += "Friend"
	fmt.Println(&s)

}
