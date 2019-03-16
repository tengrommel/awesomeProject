package main

import (
	"sync"
)

var wg sync.WaitGroup

func main() {
	// chan 的死锁
	wg.Add(2)
	cha := make(chan int)
	chb := make(chan int)
	// a
	go func() {
		<-cha
		chb <- 123
		wg.Done()
	}()
	go func() {
		<-chb
		cha <- 123
		wg.Done()
	}()
	wg.Wait()

	// 读写死锁
	var rwm sync.RWMutex
	go func() {
		rwm.RLock()
		<-cha
		rwm.RUnlock()
	}()
	go func() {
		rwm.Lock()
		cha <- 123
		rwm.Unlock()
	}()
}
