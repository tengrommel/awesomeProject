package main

import (
	"fmt"
	"time"
)

/*
协程coroutine
轻量级的线程
非抢占式多任务处理，由协程主动交出控制权
编译器/解释器/虚拟机层面的多任务
多个协程可能在一个或多个线程上运行
Subroutines are special cases of more general program components, called co routines.
In contract to the un symmetric
*/

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				fmt.Println("Hello", i)
			}
		}(i)
	}
	time.Sleep(1 * time.Millisecond)
	fmt.Println(a)
}
