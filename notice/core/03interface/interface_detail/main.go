package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type integer int

func (i integer) Say() {
	fmt.Println("integer Say i =", i)
}

type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()~~")
}

func (m Monster) Say() {
	fmt.Println("Monster Say()~~")
}

func main() {
	var stu Stu // 结构体变量 实现了Say()
	var a AInterface = stu
	a.Say()
	var i integer = 10
	var b AInterface = i
	b.Say()
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}
