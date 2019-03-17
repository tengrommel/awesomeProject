package main

import (
	"flag"
	"io"
	"log"
	"net"
	"sync"
)

var (
	target = flag.String("target", "www.baidu.com", "target host")
)

func handleConn(conn net.Conn) {
	// 建立到目标服务器的连接
	var (
		remote net.Conn
		err    error
	)
	remote, err = net.Dial("tcp", *target)
	if err != nil {
		log.Print(err)
		conn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	// go 读取conn的数据，发送到remote，知道conn的EOF，关闭remote
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	// go 接收remote的数据发送到客户端(conn)，知道remote的EOF，关闭conn
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	// 等待两个协程结束
	wg.Wait()
}

func main() {
	flag.Parse()
	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}
	// 建立listen
	for {
		conn, _ := l.Accept()
		go handleConn(conn)
		// accept new connection
	}
}
