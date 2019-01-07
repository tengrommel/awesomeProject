package main

import "fmt"

func main() {
	s := "你好吗aa"
	fmt.Printf("% x\n\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	fmt.Println(string(r))
}
