package main

import "fmt"

type integer int

func (i integer) Print() {
	fmt.Println("i=", i)
}

// 编写一个方法，可以改变i的值
func (i *integer) change() {
	*i = *i + 1
}

type Student struct {
	Name string
	Age  int
}

// 给Student实现方法String()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main() {
	var i integer = 10
	i.Print()
	i.change()
	i.Print()
	// 定义一个Student变量
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	fmt.Println(&stu)
}
