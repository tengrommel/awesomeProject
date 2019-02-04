package main

import "fmt"

type A struct {
	Name string
	age  int
}

func (a *A) SayOK() {
	fmt.Println("A SayOK", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string
}

func (b *B) SayOK() {
	fmt.Println("B sayOK", b.Name)
}

func main() {
	var b B
	b.A.Name = "tom"
	b.A.age = 19
	b.A.SayOK()
	b.A.hello()

	// 上面的写法可以简化
	b.Name = "smith"
	b.age = 20
	b.SayOK()
	b.hello()
	b.A.SayOK()
}
