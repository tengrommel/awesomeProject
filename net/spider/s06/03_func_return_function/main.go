package main

import "fmt"

func main() {
	hi := sayGreetings("ESP")
	fmt.Println(hi())
	hi = sayGreetings("GER")
	fmt.Println(sayGreetings("ENG")())
	fmt.Println()
	fmt.Printf("%T \n", hi)
	fmt.Printf("%T \n", hi())
}

func sayGreetings(lang string) func() string {
	if lang == "ESP" {
		return func() string {
			return "Hello"
		}
	} else if lang == "GER" {
		return func() string {
			return "Hallo"
		}
	}
	return func() string {
		return "OK"
	}
}
