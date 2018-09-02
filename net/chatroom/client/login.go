package main

import "fmt"

// 写一个函数， 完成登陆
func login(userId int, userPwd string) (err error) {
	// 下一步就要开始定协议...
	fmt.Printf(" userId = %d userPwd = %s", userId, userPwd)
	return nil
}
