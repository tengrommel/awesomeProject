package main

import (
	"awesomeProject/notice/core/01object/factory/model"
	"fmt"
)

func main() {
	// 创建student实例
	//var stu = model.Student{
	//	Name:  "tom",
	//	score: 78.9,
	//}
	var stu = model.NewStudent("tom~", 88.8)
	fmt.Println(stu, stu.GetScore())
}
