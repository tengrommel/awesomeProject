package main

import "fmt"

func main() {
	s := "你好啊"
	for i, v := range s {
		fmt.Printf("%#U starts at byte position %d\n", v, i)
	}

	fmt.Println([]byte(s))
	fmt.Println()
	s = "你aadfd"
	for i, v := range s {
		fmt.Printf("%#U starts at byte position %d\n", v, i)
	}
}
