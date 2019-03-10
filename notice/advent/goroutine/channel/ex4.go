package main

import (
	"fmt"
	"time"
)

/*
三个协程分别数数
要求主协程恰好在所有协程工作结束时结束
*/
func Count(grName string, n int, chanQuit chan string) {
	for i := 1; i <= n; i++ {
		fmt.Println(grName, i)
		time.Sleep(200 * time.Millisecond)
	}
	fmt.Println(grName, "工作完毕")
	chanQuit <- grName + "工作完成"
}

func main() {
	chanQuit := make(chan string, 3)
	go Count("son", 10, chanQuit)
	go Count("daughter", 7, chanQuit)
	Count("main", 5, chanQuit)
	//阻塞等待
	for i := 0; i < 3; i++ {
		x := <-chanQuit
		fmt.Println(x)
	}
	fmt.Println("over")
}
