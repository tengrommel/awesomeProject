package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "How 你 cómo お元 Wie geht 잘 How 你 có お元 Wie geht 잘 Ho có Wie"
	r := []rune(s)
	fmt.Println(len(s), utf8.RuneCountInString(s))
	counts := make(map[rune]int)
	fmt.Println()
	for i := range r {
		counts[r[i]]++
	}
	fmt.Printf("%v\n", counts)
}
