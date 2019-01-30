package main

import "fmt"

// 定义一个Cat结构体，将Cat的各个字段/属性信息，放入到Cat结构体进行管理
type Cat struct {
	Name  string
	Age   int
	Color string
	Hobby string
	Score [3]int
}

func main() {
	/*
		张老太养了20只猫猫：一只名字叫小白，今年3岁，白色。
		还有一只叫小花，今年100岁，花色。
		请编写一个程序，当用户输入小猫的名字时，就显示该猫的名字，年龄，颜色。
		如果用户输入的小猫名字错误，则显示张老太没有这只猫猫。
	*/

	/*
		// 1、使用变量的处理
		var cat1Name = "小白"
		var cat1Age = 3
		var cat1Color = "白色"
		var cat2Name = "小花"
		var cat2Age = 100
		var cat2Color = "花色"
		// 2、使用数组解决
		var catNames = [...]string{"小白", "小花"}
		var catAges = [...]int{3, 100}
		var catColors = [...]string{"白色", "花色"}

		现有技术的解决方案
		 1）使用变量或者数组来解决养猫的问题，不利于数据的管理和维护。因为名字，年龄，颜色都是属于一只猫，但是这里是分开保存。
		 2）如果我们希望对一只猫的属性（名字、年龄，颜色）进行操作（绑定方法），也不好处理
		 3）引出我们要讲解的技术-结构体

		思考步骤
		1、将一类事物的特征提取出来，形成一个新的数据类型，就是一个结构体
		2、通过这个结构体，我们可以创建多个变量（实例/对象）
		3、事物可以是猫类，也可以是Person，Fish或是某个工具类...
	*/

	// 使用struct来完成案例
	// 创建一个Cat的变量
	var cat1 Cat // var a int
	cat1.Name = "小白"
	cat1.Age = 3
	cat1.Color = "白色"
	cat1.Hobby = "吃🐟"
	fmt.Println("cat1=", cat1)
	fmt.Println("猫猫的信息如下：")
	fmt.Println("name=", cat1.Name)
	fmt.Println("age=", cat1.Age)
	fmt.Println("color=", cat1.Color)
	fmt.Println("hobby=", cat1.Hobby)
}
