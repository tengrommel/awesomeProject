package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	for i := 1; i <= 10; i++ {
		go func(i int) {
			if i == 5 {
				// 出让当前协程的资源
				runtime.Gosched()
			}
			// 协程内的代码执行在子协程
			// 协程有一定的初始化时间
			fmt.Println("子协程初始化完毕", i)
			for j := 1; j < 10; j++ {
				fmt.Printf("协程%d: %d\n", i, j)
			}
		}(i)
	}
	time.Sleep(1 * time.Second)
}
