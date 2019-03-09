package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:     20,
		MaxActive:   1000,
		IdleTimeout: time.Second * 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		}}
	defer pool.Close()
	for i := 0; i < 10; i++ {
		// 开辟一条并发协程执行happy任务
		go func(pool *redis.Pool, i int) {
			conn := pool.Get()
			defer conn.Close()
			reply, err := conn.Do("set", "conn"+strconv.Itoa(i), i)
			s, _ := redis.String(reply, err)
			fmt.Println(s)
		}(pool, i)
		go getConnFromPoolAndHappy(pool, i)
	}
	time.Sleep(3 * time.Second)
}

func getConnFromPoolAndHappy(pool *redis.Pool, i int) {
	conn := pool.Get()
	defer conn.Close()
	reply, err := conn.Do("set", "conn"+strconv.Itoa(i), i)
	s, _ := redis.String(reply, err)
	fmt.Println(s)
}
