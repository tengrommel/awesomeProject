package main

import (
	"fmt"
	"sync"
	"time"
)

/*
银行账户案例
创建银行账户类
存取款方法需要并发安全（不允许并发访问余额）
查询余额和打印流水不需要
创建账户实例，并发执行存取款、查询余额、打印流水操作
*/

type Account struct {
	money int
	// 账户的互斥锁
	mt sync.Mutex
}

// 存钱 必须保证并发安全 不允许并发操作
func (a *Account) SaveMoney(n int) {
	a.mt.Lock()
	defer a.mt.Unlock()
	fmt.Println("save money开始")
	<-time.After(3 * time.Second)
	a.money += n
	fmt.Println("save money结束")
}

// 取钱 必须保证并发安全 不允许并发操作
func (a *Account) GetMoney(n int) {
	a.mt.Lock()
	defer a.mt.Unlock()
	fmt.Println("get money开始")
	<-time.After(3 * time.Second)
	a.money -= n
	fmt.Println("get money结束")
}

// 查询余额
func (a *Account) Query() {
	fmt.Println("query money开始")
	<-time.After(3 * time.Second)
	fmt.Println("当前余额：", a.money)
	fmt.Println("query money结束")
}

func (a *Account) PrintHistory() {
	fmt.Println("打印流水开始")
	<-time.After(3 * time.Second)
	fmt.Println("打印流水")
}

func main() {
	var wg sync.WaitGroup
	account := &Account{1000, sync.Mutex{}}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.SaveMoney(100)
			wg.Done()
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.GetMoney(100)
			wg.Done()
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.Query()
			wg.Done()
		}()
	}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			account.PrintHistory()
			wg.Done()
		}()
	}
	wg.Wait()
}
