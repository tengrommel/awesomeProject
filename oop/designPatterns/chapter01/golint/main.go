package main

import (
	"fmt"
	"reflect"
)

func main() {
	var explicitString = "Hello World!"
	implicitString := "Hello World!"
	println(explicitString)
	println(implicitString)
	fmt.Println("Variable `implicitString` is of type:", reflect.TypeOf(implicitString))
	fmt.Println("Variable `explicitString` is of type:", reflect.TypeOf(explicitString))
}
