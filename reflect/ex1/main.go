package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取的传入的变量的type，kind，值
	// 1、先获取到reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("rType:", rType)
	// 2、获取reflect.Value
	rVal := reflect.ValueOf(b)
	n2 := rVal.Int() + 2
	fmt.Println("n2=", n2)
	fmt.Printf("rVal: %v rVal type: %T\n", rVal, rVal)
	// 下面我们将rVal转成interface{}
	iV := rVal.Interface()
	num2 := iV.(int)
	fmt.Println("num2:", num2)
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	// kind 类别
	fmt.Printf("rType: %v\n", rType.Kind())
	fmt.Printf("rType: %T\n", rType)
	fmt.Println("rType:", rType)
	rVal := reflect.ValueOf(b)
	fmt.Println(rVal.Kind())
	iV := rVal.Interface()
	fmt.Printf("iV= %v, iv type=%T\n", iV, iV)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	//var num = 100
	//reflectTest01(num)
	stu := Student{
		Name: "teng",
		Age:  35,
	}
	reflectTest02(stu)
}
