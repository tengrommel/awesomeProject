package main

import (
	"fmt"
	"time"
)

/*
关闭管道，无法再写入数据，但可以读出
*/
func main() {
	bools := make(chan int, 3)

	go func() {
		time.Sleep(3 * time.Second)
		x := <-bools
		fmt.Println(x)
		x = <-bools
		fmt.Println(x)
		// 一个已经关闭且读空的管道 继续读读出零值
		x, ok := <-bools
		fmt.Println(x, ok)
	}()

	bools <- 123
	bools <- 456
	close(bools)

	for {
		time.Sleep(1 * time.Second)
	}
}
