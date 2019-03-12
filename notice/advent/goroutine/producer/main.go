package main

import (
	"fmt"
	"strconv"
	"time"
)

/*
生产者每天生产一件产品
生产者每生产一件，消费者就消费一件
消费10轮后退出
*/

// 思路
/*
生产者源源不断地向【商店管道】写入产品
消费者阻塞等待从【商店管道】读出数据
消费者每读出一次数据 就向【计数管道】写入一个数据
主协程阻塞从【计数管道】读出10个数据就结束
*/

type Product struct {
	Name string
}

func main() {
	// 创建 商店管道和计数管道
	chanShop := make(chan Product, 5)
	chanCount := make(chan int, 5)
	go Produce(chanShop)
	go Consume(chanShop, chanCount)
	for i := 0; i < 10; i++ {
		<-chanCount
	}
	fmt.Println("main over")
}

func Consume(chanShop <-chan Product, chanCount chan<- int) {
	for {
		// 从商店消费
		product := <-chanShop
		fmt.Println("消费者消费了", product)
		// 报数
		chanCount <- 1
	}
}

func Produce(chanShop chan<- Product) {
	for {
		product := Product{"产品" + strconv.Itoa(time.Now().Second())}
		// 丢入商店
		chanShop <- product
		time.Sleep(1 * time.Second)
	}
}
