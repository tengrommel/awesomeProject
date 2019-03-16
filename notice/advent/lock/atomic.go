package main

import (
	"fmt"
	"sync/atomic"
)

// 原子操作只能给基本类型

func main() {
	var a int64 = 123
	value := atomic.LoadInt64(&a)
	fmt.Println(value)
}
