package main

import "fmt"

/*
Create a slice of string, including some repeating words.
Create a map that contains only the distinct words in our slice.
*/

func main() {
	s := []string{"one", "two", "three", "four", "one", "four"}
	m := make(map[string]bool)
	for i := range s {
		word := s[i]
		if !m[word] {
			m[word] = true
		}
	}
	fmt.Println(s)
	fmt.Println(m)
}
