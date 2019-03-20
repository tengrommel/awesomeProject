package main

import (
	"fmt"
	"reflect"
)

func reflect01(b interface{}) {
	rVal := reflect.ValueOf(b)
	fmt.Printf("rVal kind=%v\n", rVal.Kind())
	//rVal.SetInt(20)
	rVal.Elem().SetInt(20)
}

func main() {
	var num = 10
	reflect01(&num)
	fmt.Println(num)

}
