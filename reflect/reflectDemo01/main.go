package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(b interface{}) {
	// 通过反射获取的传入的变量的 type，kind，值
	//1、先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)
	// 2、获取到reflect.Value
	rVal := reflect.ValueOf(b)

	n2 := 2 + rVal.Int()
	fmt.Println("n2 = ", n2)
	fmt.Printf("rVal=%v  type = %T\n", rVal, rVal)

	// 下面我们将rVal转成interface{}
	iV := rVal.Interface()
	// 将interface{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func main() {
	// 1.先定义一个int
	var num int = 100
	reflectTest01(num)
}
