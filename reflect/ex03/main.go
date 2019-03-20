package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str = "tom"
	fs := reflect.ValueOf(&str)
	fs.Elem().SetString("jack")
	fmt.Printf("%v\n", str)
}
