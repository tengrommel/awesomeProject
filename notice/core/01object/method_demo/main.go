package main

import "fmt"

type Person struct {
	Name string
}

// 给A类型绑定一个方法
func (p Person) test() {
	p.Name = "Jack"
	fmt.Println("test() name=", p.Name)
}

func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}

func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}

func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果是=", res)
}

func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

type Dog struct{}

func main() {
	var p Person
	p.Name = "tom"
	p.test() // 调用方法
	fmt.Println(p.Name)
	p.speak()
	p.jisuan()
	p.jisuan2(10)
	res := p.getSum(10, 20)
	fmt.Println("res=", res)
}
