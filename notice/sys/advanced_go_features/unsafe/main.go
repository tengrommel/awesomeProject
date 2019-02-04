package main

import (
	"fmt"
	"unsafe"
)

/*
Unsafe code is Go code that bypasses the type safety and memory security of Go and requires the use of the unsafe package.
You will most likely never need to use unsafe code in your Go programs but if for some strange reason you ever need to,
it will probably have to do with pointers.
*/

func main() {
	var value int64 = 5

	var p1 = &value
	var p2 = (*int32)(unsafe.Pointer(p1))
	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 312121321321213212
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
	*p1 = 31312132
	fmt.Println(value)
	fmt.Println("*p2: ", *p2)
}
