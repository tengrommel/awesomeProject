package main

import "fmt"

type Stu struct {
	Name string
	Age  int
}

func main() {
	// 在创建结构体变量时，就直接指定字段的值
	var stu1 = Stu{"小明", 19}
	stu2 := Stu{"小明~", 20}

	fmt.Println(stu1)
	fmt.Println(stu2)

	var stu3 = &Stu{"小王", 29}
	fmt.Println(stu3)
}
