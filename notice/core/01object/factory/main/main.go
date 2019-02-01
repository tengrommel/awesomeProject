package main

import (
	"../model"
	"fmt"
)

func main() {
	// 创建student实例
	var stu = model.Student{
		Name:  "tom",
		Score: 78.9,
	}
	fmt.Println(stu)
}
