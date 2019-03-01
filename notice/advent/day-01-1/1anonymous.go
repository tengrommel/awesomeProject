package main

import (
	"fmt"
	"time"
)

var ones = []string{"关羽", "张飞", "赵云"}

func main() {
	liu := shell("刘备", 1)
	zhu := shell("诸葛", 0)
	pang := shell("帮控", 2)

	go func() {
		for i := 0; i < 70; i++ {
			kill := liu()
			fmt.Println(kill)
		}
	}()
	go func() {
		for i := 0; i < 80; i++ {
			kill := zhu()
			fmt.Println(kill)
		}
	}()
	go func() {
		for i := 0; i < 80; i++ {
			kill := pang()
			fmt.Println(kill)
		}
	}()

	time.Sleep(20 * time.Second)
}

func shell(master string, startNum int) func() string {
	startIndex := startNum
	logic := func() string {
		one := ones[startIndex]
		startIndex++
		if startIndex > 2 {
			startIndex = 0
		}
		return master + "麾下：" + one
	}
	return logic
}
