package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mt sync.Mutex

	/*
		一次只能被一个协程锁定
		其余想要抢到这把锁的协程 只能阻塞等待至前面的协程将锁释放
	*/

	var wg sync.WaitGroup
	var money = 200
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("协程%d开始抢锁\n", i)
			mt.Lock()
			fmt.Printf("协程%d抢锁成功!\n 开始操作数据\n", i)
			for j := 0; j < 100; j++ {
				money += 1
			}
			<-time.After(5 * time.Second)
			mt.Unlock()
			fmt.Printf("协程%d将锁释放!\n", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(money)
}
