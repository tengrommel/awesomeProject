package main

import (
	"fmt"
	"sync"
	"time"
)

/*
需求：现在要计算1-200的各个数的阶乘，并且把各个数的阶乘放入到map中。
最后显示出来。要求使用goroutine完成

思路
1、编写一个函数，来计算各个数的阶乘，并放入到map中
2、我们启动的协程多个，统计的将结果放入到map中
3、map应该做成全局的
*/

var (
	myMap = make(map[int]int, 10)
	// 声明全局的互斥锁
	// lock是一个全局的互斥锁
	// sync 是包：synchornized
	lock sync.Mutex
)

func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	// 加锁
	lock.Lock()
	myMap[n] = res
	// 解锁
	lock.Unlock()
}

func main() {
	for i := 1; i <= 200; i++ {
		go test(i)
	}
	// 休眠10秒
	time.Sleep(time.Second * 10)
	lock.Lock()
	for v, k := range myMap {
		fmt.Printf("map[%d]=%d\n", v, k)
	}
	lock.Unlock()
}
