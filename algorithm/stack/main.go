package main

import (
	"fmt"
	"github.com/pkg/errors"
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

// 遍历栈，注意需要从栈顶开始遍历
func (s *Stack) List() {
	// 先判断栈是否为空
	if s.Top == -1 {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	//curTop := s.Top
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1,
	}
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.List()
	val, _ := stack.Pop()
	fmt.Printf("出栈:%v\n", val)
}
