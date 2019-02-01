package main

import "fmt"

type Person struct {
	Name string
}

// 函数
func test01(p Person) {
	fmt.Println(p.Name)
}

func test02(p *Person) {
	fmt.Println(p.Name)
}

func (p Person) test03() {
	p.Name = "jack"
	fmt.Println("test03() =", p.Name)
}

func (p *Person) test04() {
	p.Name = "marry"
	fmt.Println("test04() =", p.Name)
}

func main() {
	p := Person{"tom"}
	test01(p)
	test02(&p)

	p.test03()
	fmt.Println("main() p.name =", p.Name)

	(&p).test03()                          // 依然是值拷贝 从形式上是传入地址
	fmt.Println("main() p.name =", p.Name) // tom

	(&p).test04()
	fmt.Println("main() p.name =", p.Name) // marry

	p.test04() // 等价 (&p).test04，从形式上是值传入，但本质上是地址拷贝
}
