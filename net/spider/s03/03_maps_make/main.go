package main

import "fmt"

func main() {
	days := make(map[string]int)
	//var days map[string]int
	//if days == nil {
	//	fmt.Println("'days' is nil. I'm going to allocate memory to it.")
	//	days = make(map[string]int)
	//	return
	//}
	fmt.Println(days)
	days["Sun"] = 3
	days["Sun"] = 2
	fmt.Println(days)

	days["Mon"] = 1
	days["Mon"]++
	fmt.Println(days)

	fmt.Println(days["Mon"])
	fmt.Println(days["NoDay"])
}
