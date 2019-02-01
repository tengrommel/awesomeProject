package main

import "fmt"

type MethodUtils struct{}

// 给MethodUtils编写方法

func (mu MethodUtils) Print() {
	for i := 0; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) Print2(m int, n int) {
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数...")
	} else {
		fmt.Println(num, "是奇数...")
	}
}

func (mu *MethodUtils) Print3(n int, m int, key string) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (c *Calcuator) getSum() float64 {
	return c.Num1 + c.Num2
}

func (c *Calcuator) getSub() float64 {
	return c.Num1 - c.Num2
}

func (c *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = c.Num1 + c.Num2
	case '-':
		res = c.Num1 - c.Num2
	case '*':
		res = c.Num1 * c.Num2
	case '/':
		res = c.Num1 / c.Num2
	default:
		fmt.Println("运算符输入有误...")
	}
	return res
}

func main() {
	var mu MethodUtils
	mu.Print()
	mu.Print2(5, 6)
	areaRes := mu.area(2.5, 8.7)
	fmt.Println("面积为=", areaRes)
	mu.JudgeNum(10)
	mu.JudgeNum(11)
	mu.Print3(7, 20, "@")
	var calcuator Calcuator
	calcuator.Num1 = 1.2
	calcuator.Num2 = 2.2
	fmt.Printf("sum = %v \n", fmt.Sprintf("%.2f", calcuator.getSum()))
	fmt.Printf("sub = %v \n", fmt.Sprintf("%.2f", calcuator.getSub()))
	res := calcuator.getRes('-')
	fmt.Printf("res =%v\n", fmt.Sprintf("%.2f", res))
}
