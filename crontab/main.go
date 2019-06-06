package main

import (
	"fmt"
	"strings"
)

func main() {
	stringList := []string{"abc", "bdffa","fdafd"}
	fmt.Println(strings.Join(stringList, ","))
	stringsEx := "this&ff"
	sList := strings.Split(stringsEx,"&")
	for i := range sList{
		fmt.Println(sList[i])
	}
}
