package main

import (
	"fmt"
	"time"
)

// 终止定时器

func main() {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("计时开始", time.Now())

	go func() {
		time.Sleep(3 * time.Second)
		//ok := timer.Stop()
		timer.Reset(1 * time.Second)
		//if ok {
		//	fmt.Println("炸弹已拆除")
		//	os.Exit(0)
		//}
	}()
	endTime := <-timer.C
	fmt.Println("炸弹引爆于", endTime)
}
