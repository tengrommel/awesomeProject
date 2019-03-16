package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

// 控制并发数为5

func GetSqrt(name string, n int, chSem chan string) {
	// 想执行先注册
	// 能写入就执行，写不进去就阻塞到能写入为止
	chSem <- name
	ret := math.Sqrt(float64(n))
	time.Sleep(time.Second)
	fmt.Printf("%s的平方根是%.2f\n", name, ret)
	// 任务结束完毕，从信号量控制管道注销自己，以便为其它并发任务腾出空间
	<-chSem
}

func main() {
	/*
		并发数（信号量）控制管道
		凡要并发执行的协程必须先将协程的名字注册到该管道
	*/
	chSem := make(chan string, 5)

	for i := 0; i < 100; i++ {
		go GetSqrt("协程"+strconv.Itoa(i), i, chSem)
	}
	for {
		time.Sleep(time.Second)
	}
}
