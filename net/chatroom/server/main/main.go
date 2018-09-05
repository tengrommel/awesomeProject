package main

import (
	"awesomeProject/net/chatroom/server/model"
	"fmt"
	"net"
	"time"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 这里需要延时关闭
	defer conn.Close()
	// 这里调用主控，创建一个
	process := Processor{
		Conn: conn,
	}
	err := process.process2()
	if err != nil {
		fmt.Println("客户端和服务器通讯错误=err", err)
		return
	}
}

// 这里我们编写一个函数，完成对UserDao的初始化任务
func initUserDao() {
	// 这里的pool本身就是全局变量
	// 这里要注意一个初始化的顺序问题
	// initPool 再 initUserPool
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	// 当服务器启东时，我们就去初始化我们的redis的连接池
	initPool("localhost:6379", 16, 0, 300*time.Second)
	initUserDao()
	// 提示信息
	fmt.Println("服务器[新的结构]在8889端口监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	// 一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来链接服务器")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}
		// 一旦链接成功，则启动一个协程和客户端保持通讯
		go process(conn)
	}
}
