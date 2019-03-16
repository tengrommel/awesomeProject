package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var rwm sync.RWMutex
	//// 锁定为写模式 -- 一路只写 不可读
	//rwm.Lock()
	//// 解锁写模式
	//rwm.Unlock()
	//// 锁定为读模式 -- 多路只读
	//rwm.RLock()
	//// 释放读写锁
	//rwm.RUnlock()

	var wg sync.WaitGroup
	var rwm sync.RWMutex
	/*
		数据库的一写多读
	*/
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			rwm.RLock()
			defer wg.Done()
			defer rwm.RUnlock()
			fmt.Println("读取数据库")
			<-time.After(3 * time.Second)
		}()
		wg.Add(1)
		go func() {
			rwm.Lock()
			defer wg.Done()
			defer rwm.Unlock()
			fmt.Println("写入数据库")
			<-time.After(3 * time.Second)
		}()
	}
	wg.Wait()
}
