package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("当前协程的可用CPU逻辑核心数为：", runtime.NumCPU())
	// max processors
	// 将可用CUP逻辑核心数设置为4，并返回先前的配置
	fmt.Println(runtime.GOMAXPROCS(4))
	fmt.Println(runtime.GOMAXPROCS(2))
	fmt.Println(runtime.GOMAXPROCS(2))

}
