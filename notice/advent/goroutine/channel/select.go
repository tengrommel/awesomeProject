package main

import (
	"fmt"
	"time"
)

func main() {
	chA := make(chan int, 5)
	chB := make(chan int, 4)
	chB <- 123
	chB <- 123
	chB <- 123
	chB <- 123
	chC := make(chan int, 3)
	chC <- 123
	chC <- 123
	chC <- 123
	//go TaskA(chA)
	//go TaskB(chB)
	//go TaskC(chC)
OUTER:
	for {
		select {
		case chA <- 123:
			fmt.Println("执行任务A")
			time.Sleep(time.Second)
		case <-chB:
			fmt.Println("执行任务B")
			time.Sleep(time.Second)
		case <-chC:
			fmt.Println("执行任务C")
			time.Sleep(time.Second)
		default:
			fmt.Println("全部任务已结束")
			break OUTER
		}
	}
}

func TaskA(ch chan int) {
	for {
		fmt.Println("TaskA")
		ch <- 123
		time.Sleep(time.Second)
	}
}

func TaskB(ch chan int) {
	for {
		fmt.Println("TaskB")
		<-ch
		time.Sleep(time.Second)
	}
}

func TaskC(ch chan int) {
	for {
		fmt.Println("TaskC")
		<-ch
		time.Sleep(time.Second)
	}
}
