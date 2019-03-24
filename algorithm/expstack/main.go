package main

import (
	"errors"
	"fmt"
	"strconv"
)

type Stack struct {
	MaxTop int    // 表示我们栈最大可以存放个数
	Top    int    // 表示栈顶
	arr    [5]int // 数组模拟栈
}

func (s *Stack) Push(val int) (err error) {
	if s.Top == s.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("Stack full")
	}
	s.Top++
	// 放入数据
	s.arr[s.Top] = val
	return
}

func (s *Stack) Pop() (val int, err error) {
	if s.Top == -1 {
		fmt.Println("stack empty!")
		return 0, errors.New("stack empty")
	}
	// 先取值，再s.Top--
	val = s.arr[s.Top]
	s.Top--
	return val, nil
}

// 判断一个字符是不是一个运算符[+,-,*,/]
func (s *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算的方法
func (s *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num2 * num2
	case 43:
		res = num2 + num1
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("运算符错误")
	}
	return res
}

// 编写一个方法，返回某个运算符的优先级【程序员定义】
// [* / => 1 + - => 0]
func (s *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 {
		res = 1
	} else if oper == 43 || oper == 45 {
		res = 0
	}
	return res
}

func main() {
	// 数栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	// 符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "3+2*6-2"
	// 定义一个index，帮助扫描exp
	index := 0
	// 为了配合运算，我们定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	for {
		ch := exp[index : index+1]
		temp := int([]byte(ch)[0])  //就是字符对应的ASCII码
		if operStack.IsOper(temp) { // 说明是符号
			// 两个逻辑
			// 如果是空栈，直接入栈
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					numStack.Push(result)
					// 当前的符号压入符号栈
					operStack.Push(temp)
				} else {
					operStack.Push(temp)
				}
			}
		} else { // 说明是数
			val, _ := strconv.ParseInt(ch, 10, 64)
			numStack.Push(int(val))
		}
		// 继续扫描
		// 先判断index是否已经扫描到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++
	}
	for {
		if operStack.Top == -1 {
			break // 退出条件
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		numStack.Push(result)
	}
	// 如果代码没有问题，表达式正确则结果就是numStack中唯一的数
	res, _ := numStack.Pop()
	fmt.Println(res)

}
