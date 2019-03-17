package main

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

// 1、握手
// 2、获取客户端代理的请求
// 3、开始代理
func readAddr(r *bufio.Reader) (string, error) {
	version, _ := r.ReadByte()
	log.Printf("verison: %d", version)
	if version != 5 {
		return "", errors.New("bad verison")
	}
	cmd, _ := r.ReadByte()
	log.Printf("cmd: %d", cmd)
	if cmd != 1 {
		return "", errors.New("bad cmd")
	}
	// skip
	r.ReadByte()
	addrType, _ := r.ReadByte()
	log.Printf("%d", addrType)
	if addrType != 3 {
		return "", errors.New("bad addr type")
	}
	// 读取一个字节的数据，代表后面紧跟着的域名的长度
	// 读取n个字节得到域名，n根据上一步得到的结果来决定
	// addrLen
	// addr
	addrLen, _ := r.ReadByte()
	addr := make([]byte, addrLen)
	io.ReadFull(r, addr)
	log.Printf("%v", addr)
	var port int16
	binary.Read(r, binary.BigEndian, &port)
	return fmt.Sprintf("%s:%d", addr, port), nil
}

func handleShake(r *bufio.Reader, conn net.Conn) error {
	version, _ := r.ReadByte()
	log.Printf("version: %d", version)
	if version != 5 {
		return errors.New("bad version")
	}
	nMethods, _ := r.ReadByte()
	log.Printf("nMethods d%", nMethods)

	buf := make([]byte, nMethods)
	io.ReadFull(r, buf)
	log.Printf("%v", buf)

	resp := []byte{5, 0}
	conn.Write(resp)
	return nil
}

func handle(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	handleShake(r, conn)
	log.Printf("version: %d", version)
	addr, _ := readAddr(r)
	log.Print(addr)
	resp := []byte{}
	conn.Write(resp)
}

func main() {
	l, err := net.Listen("tcp", ":8021")
	if err != nil {
		log.Fatal(err)
	}
	// 建立listen
	for {
		conn, _ := l.Accept()
		go handle(conn)
		// accept new connection
	}
}
