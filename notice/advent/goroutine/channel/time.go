package main

import (
	"fmt"
	"time"
)

/*
三秒钟之后宣布重大决定
*/

func main() {
	timer := time.NewTicker(3 * time.Second)
	fmt.Println("计数开始", time.Now())
	x := <-timer.C
	fmt.Println("引爆于", x)
	fmt.Println("我不信你")
}
