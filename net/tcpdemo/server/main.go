package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	// 这里我们循环的接受客户端发送的数据
	defer conn.Close() // 关闭conn
	for {
		// 创建一个新的切片
		buf := make([]byte, 1024)
		// conn.Read(buf)
		// 1.等待客户端通过conn发送信息
		// 2.如果客户端没有write[发送]，那么协程就阻塞在这里
		fmt.Printf("服务器在等待客户端%s发送信息\n", conn.RemoteAddr().String())
		n, err := conn.Read(buf) // 从conn读取
		if err == io.EOF {
			// io.EOF
			fmt.Println("远程服务器已退出 err=", err)
			return
		}
		// 3.显示客户端发送的内容到服务器显示终端
		fmt.Print(string(buf[:n]))
	}
}

func main() {
	fmt.Println("服务器开始监听了....")
	// 1.tcp表示使用网络协议是tcp
	// 2.0.0.0.0:8888表示在本地监听 8888端口
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err =", err)
		return
	}
	defer listen.Close() // 延时关闭listen
	// 循环等待客户端来连接我
	for {
		// 等待客户端来连接我
		fmt.Println("等待客户端来链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accept() err =", err)
		} else {
			fmt.Printf("Accept() suc con = 客户端ip:%v\n", conn.RemoteAddr())
		}
		// 这里起一个协程，为客户端服务
		go process(conn)
	}
	//fmt.Printf("listen=%v\n", listen)
}
