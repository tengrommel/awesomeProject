package main

import (
	"fmt"
	"sync"
	"time"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	var wg sync.WaitGroup
	fmt.Println(wg)
	wg.Add(3)
	// 向等待组当中加一个协程（注册）

	// 从等待组中减掉一个协程（注销）
	// 阻塞等待至等待组中的协程数归零

	/*
		分别使用Ticker和Timer创建两个耗时的协程A, B
		将A,B两个协程加入等待组
		A,B结束时从等待组注销
		主协程阻塞至等待组内协程数归零
	*/
	go func() {
		fmt.Println("协程A开始工作...")
		time.Sleep(3 * time.Second)
		fmt.Println("协程A over!")
		wg.Done()
	}()

	go func() {
		fmt.Println("协程B开始工作...")
		<-time.After(5 * time.Second)
		fmt.Println("协程B over!")
		wg.Done()
	}()

	go func() {
		fmt.Println("协程C开始工作...")
		ticker := time.NewTicker(1 * time.Second)
		for i := 0; i < 4; i++ {
			<-ticker.C
		}
		ticker.Stop()
		fmt.Println("协程C over!")
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("main over")
}
