package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理环形队列

type CircleQueue struct {
	maxSize int    // 4
	array   [4]int // 数组
	head    int    // 指向队列队首
	tail    int    // 指向队列的队尾
}

// 入队列 AddQueue(push) GetQueue(pop)
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	// 分析出this.tail在队列尾部，但是不包含最后的元素
	this.array[this.tail] = val // 把值给尾部
	this.tail++
	return
}

// 出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	// 取出，head是指向队首，并且含队首元素
	val = this.array[this.head]
	this.head++
	return
}

// 显示队列
func (this *CircleQueue) ListQueue() {
	// 取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}
	// 设计一个辅助的变量，指向head
	tempHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\t", tempHead, this.array[tempHead])
		tempHead = (tempHead + 1) % this.maxSize
	}
	fmt.Println()
}

// 判断环形队列为满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 判断环形队列是空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 取出环形队列有多少个元素
func (this *CircleQueue) Size() int {
	// 这是一个关键的算法。
	return (this.tail + this.maxSize - this.head) % this.maxSize
}

func main() {

	var key string
	var val int
	for {
		fmt.Println("1.输入add表示添加数据到队列")
		fmt.Println("2.输入get表示从队列获取数据")
		fmt.Println("3.输入show表示显示死队列")
		fmt.Println("4.输入text表示显示队列")
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队列数")
			fmt.Scanln(&val)
			//err := queue.AddQueue(val)
			//if err != nil {
			//	fmt.Println("加入队列ok")
			//} else {
			//	fmt.Println(err.Error())
			//}
		case "get":
			//val, err := queue.GetQueue()
			//if err != nil {
			//	fmt.Println(err.Error())
			//} else {
			//	fmt.Println("从队列中取出了一个数=", val)
			//}
		case "show":
			//queue.showQueue()
		case "exit":
			os.Exit(0)
		}
	}
}
