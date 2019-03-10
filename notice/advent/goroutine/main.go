package main

import (
	"fmt"
	"time"
)

/*
coorperate 合作
routine 日程，事务
coroutine 协程（可以相互协作的并发任务）
goroutine go语言协程
*/

func main() {
	go CountNumber()
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}

	// 同步的，阻塞的任务
	CountNumber()

	fmt.Println("main over")
}

func CountNumber() {
	for i := 91; i <= 100; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}
