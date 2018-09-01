package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func reflectTest01(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	rVal := reflect.ValueOf(b)
	iV := rVal.Interface()
	fmt.Printf("iV = %v type = %T", iV, iV)
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Println("rType=", rType)
	rVal := reflect.ValueOf(b)

	kind1 := rVal.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind = %v kind=%v\n", kind1, kind2)

	iV := rVal.Interface()
	fmt.Printf("iV = %v type = %T \n", iV, iV)
	// 将interface{}通过断言转成需要的类型
	// 这里我们就简单使用了一带检测的类型断言

	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

func main() {
	// 定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  42,
	}
	reflectTest02(stu)
}
