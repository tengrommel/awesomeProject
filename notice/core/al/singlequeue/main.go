package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体来管理队列
type Queue struct {
	maxSize int
	array   [5]int // 数组
	front   int    // 表示指向队列
	rear    int    // 表示指向队列的尾部
}

func (q *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if q.rear == q.maxSize-1 {
		return errors.New("队列已满")
	}
	q.rear++
	return nil
}

// 显示队列，找到队首，然后到遍历到队尾
func (q *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
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
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列中获取数据")
		fmt.Println("3.输入show表示显示队列")
		fmt.Println("4.输入exit表示退出")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列的数")
			fmt.Scanln(&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列ok")
			}
		case "get":
			fmt.Println("get")
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
