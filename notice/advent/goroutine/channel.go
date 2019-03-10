package main

import "fmt"

func main() {
	var myChan chan int
	fmt.Println(myChan)
	// 管道就是一片联系的内存 创建缓存能力为0的管道
	myChan = make(chan int, 1)
	// 缓存能力为0
	myChan <- 123
	// 管道已满，写不进去，主协程死锁
	myChan <- 124
	x := <-myChan
	fmt.Println(x)
	fmt.Println("main over")
}
