package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 0; i < 100000; i++ {
		intChan <- i
	}

	// 关闭intChan
	close(intChan)
}

// 从intChan取出数据，并判断是否为素数，
// 如果是，就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用for循环
	var num int
	var flag bool
	for {
		var ok bool
		num, ok = <-intChan
		if !ok { // 管道取不到
			break
		}
		// 判断num是不是素数
		flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明该num不是素数
				flag = false
				break
			}
		}
		if flag {
			// 将这个数放入到primeChan
			primeChan <- num
		}
	}
	fmt.Println("有一个primeNum协程因为取不到数据，退出")
	// 这里我们还不能关闭primeChan
	// 向exitChan写入
	exitChan <- true
}

func main() {
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000) // 放入结果
	// 标示退出的管道
	exitChan := make(chan bool, 4) // 4个

	start := time.Now().Unix()
	// 开启一个协程，向intChan放入1-8000个数
	go putNum(intChan)

	// 开启4个协程，从intChan取出数据，并判断是否为素数，如果是，就放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}
	// 这里我们主线程，进行处理
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()
		fmt.Println("耗时：", end-start)
		// 当我们从exitChan取出4个结果，就可以放心的关闭 primeNum
		close(primeChan)
	}()

	// 遍历我们的primeNum，把结果取出
	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		fmt.Printf("素数=%d\n", res)
	}
}
