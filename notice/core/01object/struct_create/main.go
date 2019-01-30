package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	//方式1

	//方式2
	p2 := Person{"mary", 20}
	//p2.Name = "Tom"
	//p2.Age = 18
	fmt.Println(p2)
	var p3 *Person = new(Person)
	// 也可以 p3.Name = "smith"
	// 原因go的设计者为了程序员使用方便，底层会对p3.Name="smith" 进行处理会给p3加上取值运算 (*p3).Name = "smith"
	(*p3).Name = "smith"
	p3.Name = "john"
	(*p3).Age = 30
	fmt.Println(*p3)

	var person *Person = &Person{"mary", 60}
	fmt.Println(*person)
}
