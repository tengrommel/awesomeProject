package main

import (
	"fmt"
	"strings"
)

func aa(str []string) {
	str[0] = "ok"
}

func bb(str []string) []string {
	// 想修改
	str = append(str, "dd")
	return str
}

func main() {
	listString := []string{"1", "2", "3"}
	fmt.Println(strings.Join(listString[:len(listString)-1], "/"))
}
