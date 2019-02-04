package main

import "fmt"

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

// 如果我们需要实现AInterface，就需要将BInterface CInterface的方法都实现
type Stu struct {
}

func (stu Stu) test01() {

}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

type T interface{}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()
	var t T = stu
	fmt.Println(t)
	var t2 interface{} = stu
	fmt.Println(t2)
}
