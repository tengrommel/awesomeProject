package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Hello你好"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
}
