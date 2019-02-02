package main

import "fmt"

// 编写一个学生考试系统
type Student struct {
	Name  string
	Age   int
	Score int
}

// 将Pupil和Graduate共有的方法也绑定到*Student
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n", stu.Name, stu.Age, stu.Score)
}

func (stu *Student) SetScore(score int) {
	stu.Score = score
}

// 给*Student增加一个方法 pupil 和 graduate都可以使用该方法
func (stu *Student) GetSum(n1 int, n2 int) int {
	return n1 + n2
}

// 编写学生考试系统
// 小学生 嵌入了一个匿名结构体
type Pupil struct {
	Student
}

func (p *Pupil) testing() {
	fmt.Println("小学生正在考试中...")
}

// 大学生
type Graduate struct {
	Student
}

func (g *Graduate) testing() {
	fmt.Println("大学生正在考试中...")
}

// 代码冗余... 高中生...
func main() {
	// 当我们对结构体嵌入了匿名结构体使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "Tom"
	pupil.Student.Age = 8
	pupil.testing()
	pupil.Student.SetScore(90)
	pupil.Student.ShowInfo()
	fmt.Println("res =", pupil.GetSum(10, 30))

	graduate := &Graduate{}
	graduate.Student.Name = "Marry"
	graduate.Student.Age = 20
	graduate.testing()
	graduate.Student.SetScore(80)
	graduate.Student.ShowInfo()
	fmt.Println("res =", pupil.GetSum(40, 50))
}
