package main

import (
	"../model"
	"fmt"
)

func main() {
	// 创建一个account变量
	account := model.NewAccount("jzh111111", "999999", 40)
	if account != nil {
		fmt.Println("创建成功=", account)
	} else {
		fmt.Println("创建失败")
	}
}
