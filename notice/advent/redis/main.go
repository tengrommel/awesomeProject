package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
)

func HandleError(err error, when string) {
	if err != nil {
		fmt.Println(when, err)
		os.Exit(1)
	}
}

func main() {
	conn, e := redis.Dial("tcp", "127.0.0.1:6379")
	HandleError(e, "redis.Dial")
	defer conn.Close()
	//reply, err := conn.Do("SET", "name", "你妹")
	reply, err := conn.Do("zrange", "mz", 0, -1)
	HandleError(err, "conn.Do Get")
	fmt.Printf("type=%T,value=%v\n", reply, reply)
	nameStr, _ := redis.Strings(reply, err)
	fmt.Println(nameStr)
}
