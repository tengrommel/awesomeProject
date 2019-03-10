package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	myChan := make(chan string, 3)
	go func() {
		for i := 0; i < 5; i++ {
			// 已经读空时，没有人写，就读出阻塞
			x := <-myChan
			fmt.Println("读出：", x)
		}
	}()
	for i := 0; i < 5; i++ {
		// 缓存已满时，没有人读，就写入阻塞
		myChan <- strconv.Itoa(i)
		fmt.Println(i, "已写入")
	}
	for {
		time.Sleep(1 * time.Second)
	}
}
