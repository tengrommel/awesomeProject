package main

import "fmt"

func main() {
	intChan := make(chan int, 3)
	intChan <- 100
	intChan <- 200
	close(intChan) // close
	// 这时不能够再写入数据到chan
	fmt.Println("okok~")
	// 当管道关闭后，读取数据是可以的
	fmt.Println(<-intChan)

	// 遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 // 放入100个数据到管道
	}

	// 如果channel没有关闭，会出现死锁现象，但可以取出数据
	// 遍历管道不能使用普通的循环结构
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
