package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err=", err)
		return
	}
	reader := bufio.NewReader(os.Stdin) // os.Stdin 代表标准输入【终端】
	// 从终端读取一行用户输入，并准备发送给服务器
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readString err=", err)
			break
		}
		// 如果用户输入的是exit就退出
		line = strings.Trim(line, "\r\n")
		fmt.Println(line == "exit")
		if line == "exit" {
			fmt.Println("客户端退出...")
			break
		}
		// 再将line发送给 服务器
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn write", err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据\n", n)
	}
}
