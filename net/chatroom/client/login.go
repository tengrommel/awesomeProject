package main

import (
	"awesomeProject/net/chatroom/common/message"
	"awesomeProject/net/chatroom/util"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 写一个函数， 完成登陆
func login(userId int, userPwd string) (err error) {
	// 下一步就要开始定协议...
	//fmt.Printf(" userId = %d userPwd = %s", userId, userPwd)
	//return nil

	// 1、连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	// 延时关闭
	defer conn.Close()

	// 2.准备通过conn发送消息给服务器
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3.创建一个LoginMes 结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	// 4、将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 5、把data赋给了message的Data字段
	mes.Data = string(data)

	// 6、将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	// 7、到这个时候data就是我们要发送的消息
	// 7.1先把data的长度发送给服务器
	// 先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// 发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	//fmt.Println("客户端，发送消息的长度", len(data))
	// 发送消息
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	// 这里还需要处理服务器端返回的消息.
	mes, err = util.ReadPkg(conn) // mes就是
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
	}
	// 将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		fmt.Println("登陆成功")
	} else if loginResMes.Code == 500 {
		fmt.Println(loginResMes.Error)
	}
	return
}
