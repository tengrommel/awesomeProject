package main

import "fmt"

/*
编写一个Student结构体，包含name、gender、age、id、score字段，
分别为string、string、结构体中声明一个say方法，
返回string类型，方法返回信息中包含所有字段值。
在main方法中，创建Student结构体实例（变量），并访问say方法，并将调用结果打印输出。
*/

type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {
	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v], age=[%v] id=[%v] score=[%v]",
		student.name, student.gender, student.age, student.id, student.score)
	return infoStr
}

type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

type Visitor struct {
	Name string
	Age  int
}

func (v *Visitor) showPrice() {
	if v.Age > 90 || v.Age <= 8 {
		fmt.Println("考虑到安全，就不要玩了")
		return
	}
	if v.Age > 18 {
		fmt.Printf("游客的名字为 %v 年龄为 %v 收费 20元\n", v.Name, v.Age)
	} else {
		fmt.Printf("游客的名字为 %v 年龄为 %v 免费\n", v.Name, v.Age)
	}
}

func main() {
	// 测试
	// 创建一个Student实例变量
	var stu = Student{
		name:   "tom",
		gender: "male",
		age:    18,
		id:     1000,
		score:  99.98,
	}
	fmt.Println(stu.say())
	// 测试代码
	var box Box
	box.len = 1.1
	box.width = 2.0
	box.height = 3.0
	volumn := box.getVolumn()
	fmt.Printf("体积为 = %.2f", volumn)
	// 测试
	var v Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
