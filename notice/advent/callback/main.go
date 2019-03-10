package main

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("hello")
}

func main() {
	go func(f func()) {
		for i := 11; i < 20; i++ {
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
		f()
	}(SayHello)
	//for i:=0;i<10;i++ {
	//	fmt.Println(i)
	//	time.Sleep(1 *time.Second)
	//}
	time.Sleep(10 * time.Second)
}
