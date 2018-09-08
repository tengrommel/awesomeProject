package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	maxSize int
	array   [5]int // 数组
	front   int    // 表示指向队列的最前面
	rear    int    // 表示指向队列的尾部
}

// 添加数据到队列
func (this *Queue) AddQueue(val int) (err error) {
	// 先判断队列已满
	if this.rear == this.maxSize-1 { // 重要的提示；rear是包含队列尾部
		return errors.New("queue full")
	}
	this.rear++ // 将rear后移
	this.array[this.rear] = val
	return
}

// 显示队列，找到队首然后遍历到队尾
func (this *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
}

// 编写一个主函数测试，测试
func main() {
	// 先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
	}
	var key string
	var val int
	for {
		fmt.Println("1、输入add表示添加数据到队列")
		fmt.Println("2、输入get表示从队列中获取数据")
		fmt.Println("3、输入show表示显示队列")
		fmt.Println("4、输入exit表示显示队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("加入队列ok")
			} else {
				fmt.Println(err.Error())
			}
		case "get":
			fmt.Println("get")
		}
	}
}
