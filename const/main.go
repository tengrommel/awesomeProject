package main

import "fmt"

func main() {
	var num int
	// 常量在声明的时候必须给值
	// 常量只能修饰bool 数值 string
	const tax int = 0
	// 常量是不能修改的
	fmt.Println(num, tax)
	const (
		a = iota
		b = iota
		c
		d = iota
	)
	fmt.Println(a, b, c, d)
}
