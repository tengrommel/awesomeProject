package main

import "fmt"

func main() {
	myChan := make(chan int, 5)
	myChan <- 123
	fmt.Println("管道的长度是", len(myChan))
	fmt.Println("管道的容量是", cap(myChan))
	// 已满的管道无法再进行写入 已读空的管道无法再读出
	myChan <- 122
	myChan <- 345
	myChan <- 123
	myChan <- 124
	myChan <- 126
}
