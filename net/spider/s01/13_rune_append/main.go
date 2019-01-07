package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "byt 再见" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	fmt.Printf("%q\n", []rune("bye 再见"))
}
