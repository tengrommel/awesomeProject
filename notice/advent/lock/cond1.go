package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	bitcoinRising := false
	cond := sync.NewCond(&sync.Mutex{})

	go func() {
		for {
			cond.L.Lock()
			bitcoinRising = true
			cond.Broadcast() // 向多个 向一个signal
			cond.L.Unlock()

			time.Sleep(1 * time.Second)
			bitcoinRising = false
			time.Sleep(3 * time.Second)
		}

	}()

	for {
		cond.L.Lock()
		for !bitcoinRising {
			fmt.Println("比特币没涨")
			cond.Wait()
		}
		fmt.Println("开始投资比特币")
		cond.L.Unlock()
	}
}
