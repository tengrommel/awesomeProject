package main

import (
	"errors"
	"fmt"
)

// 使用一个结构体来管理队列

type Queue struct {
	maxSize int
	array   [maxSize]int // 数组
	front   int          // 表示指向队列
	rear    int          // 表示指向队列的尾部
}

func (q *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if q.rear == q.maxSize-1 {
		return errors.New("队列已满")
	}
	q.rear++
	return nil
}

func (q *Queue) ShowQueue() {
	fmt.Println("队列当前的情况是：")
	for i := q.front + 1; i <= q.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, q.array[i])
	}
}

func main() {

}
