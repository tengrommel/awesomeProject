package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b)
	fmt.Println(a, b)
	// 1、创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇～"}

	// 2、将monster变量序列化为json格式字符串
	jsonMonster, err := json.Marshal(monster)
	if err != nil {
		log.Fatal("json 处理错误 ", err)
	}
	fmt.Println("jsonMonster", string(jsonMonster))
}
