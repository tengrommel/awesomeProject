package main

import (
	"fmt"
	"time"
)

func DoTask(no int) {
	for i := 1; i < 10; i++ {
		fmt.Println(no)
		time.Sleep(500 * time.Millisecond)
	}
}

// 主死从随
func main() {
	for i := 1; i < 100; i++ {
		go DoTask(i)
	}
	time.Sleep(6 * time.Second)
	// 主协程一死 子协程都随着死
	fmt.Println("Main over")
}
