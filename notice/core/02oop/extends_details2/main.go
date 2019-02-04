package main

import "fmt"

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
	//Name string
}

type D struct {
	a A // 有名结构体
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int
	n int
}

func main() {
	var c C
	// 如果C 没有Name字段 而A和B有Name 这是就必须通过指定匿名结构体名字来区分
	// 所以c.Name 就会爆编译错误
	//c.Name = "tom" error
	c.A.Name = "tom"
	fmt.Println("c")
	var d D
	d.a.Name = "jack"
	tv1 := TV{Goods{"电视机001", 5000.99}, Brand{"海尔", "山东"}}
	tv2 := TV{
		Goods{
			Name:  "电视机002",
			Price: 5000.99,
		},
		Brand{
			Name:    "夏普",
			Address: "东京"}}
	tv3 := TV2{
		&Goods{"电视机003", 7000.99},
		&Brand{"创维", "河南"},
	}
	fmt.Println("tv1:", tv1)
	fmt.Println("tv2", tv2)
	fmt.Println("tv3 Goods", *tv3.Goods)
	//演示一下匿名字段是基本数据类型的使用
	var e E
	e.Name = "狐狸精"
	e.Age = 300
	e.int = 20
	e.n = 33
	fmt.Println("E =", e)
}
