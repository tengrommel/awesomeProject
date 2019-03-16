package main

import (
	"fmt"
	"sync"
	"time"
)

/*
服务器负载控制
- 监听最大客户端连接数
- 服务端协程：只要服务器客户端数量达到阀值
- 监察协程：消减客户端数量后通知服务端（再次接入）
*/

const MAX_CONNECTIONS = 3

type Server struct {
	connections int
	chAlarm     chan bool
	cond        *sync.Cond
}

func (s *Server) Server() {
	s.cond = sync.NewCond(&sync.Mutex{})
	s.chAlarm = make(chan bool)
	go func() {
		for {
			// 阻塞监听是否有预警
			<-s.chAlarm
			// 削减客户端数量
			s.cond.L.Lock()
			<-time.After(3 * time.Second)
			s.connections--
			fmt.Println("s.connection: ", s.connections)
			s.cond.Broadcast()
			s.cond.L.Unlock()
			// 发送预警解除广播
		}
	}()

	for {
		s.cond.L.Lock()
		for s.connections == MAX_CONNECTIONS {
			s.chAlarm <- true
			fmt.Println("过载已解除")
			s.cond.Wait()
		}
		time.Sleep(1 * time.Second)
		s.connections++
		fmt.Println("已接入客户端: s.connections=", s.connections)
		s.cond.L.Unlock()
	}
}

func main() {
	new(Server).Server()
}
