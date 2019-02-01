package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// 为了提高效率，通常我们方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	fmt.Printf("c 是 *Circle 指向地址 = %p\n", c)
	c.radius = 10
	return math.Pi * c.radius * c.radius
}

func main() {
	//var c Circle
	//c.radius = 4.0
	//res := c.area()
	//fmt.Println("面积是=", res)

	var c Circle
	fmt.Printf("main c 结构体变量地址 = %p\n", &c)
	c.radius = 5.0
	// 因为c是指针标准访问字段的方式如下
	//res2 := (&c).area2()
	res2 := c.area2()
	// 编译器底层做了优化 (&c).area2() 等价 c.area2()
	// 因为编译器会自动加上&
	fmt.Println("面积=", res2)
	// 因为是指针 所以会改变
	fmt.Println("c.radius = ", c.radius)
}
