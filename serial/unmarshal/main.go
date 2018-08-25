package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
}

// 演示将json字符串，反序列化成struct
func unmarshalstruct() {
	// 说明str在项目开发中，是通过网络传输获取到..或者是读取文件获取到
	str := "{\"Name\":\"牛魔王\",\"Age\":500,\"Birthday\":\"2011-11-11\",\"Sal\":8000,\"Skill\":\"牛魔拳\"}"
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化 monster=%v\n monster.Name=%v", monster, monster.Name)
}

// 演示将json字符串，反序列化成map
func unmarshalMap() {
	str := "{\"address\":\"洪崖洞\",\"age\":30,\"name\":\"红孩儿\"}"
	// 定义一个map
	var a map[string]interface{}
	// 反序列化
	// 注意：反序列化map，不需要make，因为make操作被封装到Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 a=%v\n", a)
}

// 演示将json字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"北京\",\"age\":\"7\",\"name\":\"jack\"}," +
		"{\"address\":[\"墨西哥\", \"夏威夷\"],\"age\":\"20\",\"name\":\"tom\"}]"
	var slice []map[string]interface{}
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 slice=%v\n", slice)
}

func main() {
	unmarshalstruct()
	unmarshalMap()
	unmarshalSlice()
}
