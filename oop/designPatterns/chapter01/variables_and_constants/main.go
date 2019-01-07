package main

import (
	"fmt"
	"reflect"
)

func main() {
	var explicit = "Hello, I'm a explicitly declared variable"
	inferred := "Hello World!"
	fmt.Println("Variable 'explicit' is of type:", reflect.TypeOf(explicit))
	fmt.Println("Variable 'inferred' is of type:", reflect.TypeOf(inferred))
}
