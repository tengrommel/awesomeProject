package main

import (
	"fmt"
	"net"
)

// 处理和客户端的通讯
func process(conn net.Conn) {
	// 这里需要延时关闭
	defer conn.Close()
	// 循环读取客户端发送的信息
	buf := make([]byte, 8096)
	for {
		fmt.Println("读取客户端发送的数据...")
		_, err := conn.Read(buf[:4])
		if err != nil {
			fmt.Println("conn.Read err=", err)
			return
		}
		fmt.Println("读到的buf=", buf[:4])
	}
}

func main() {
	// 提示信息
	fmt.Println("服务器在8889端口监听...")
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
