package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
从开源接口showapi.com中获取成语数据
将其重构为Go的结构体数据
实现在命令行进行成语释义查询功能
要求实现数据持久化
*/

type Idiom struct {
	Title      string `json:"title"`
	Spell      string `json:"spell"`
	Content    string `json:"content"`
	Sample     string `json:"sample"`
	Derivation string `json:"derivation"`
}

func main() {
	jsonStr, _ := GetJson("http://route.showapi.com/1196-1?showapi_appid=19988&showapi_sign=968ad4fcc2144e41b5c366838d1b0ec4&keyword=%E9%91%B2?page=1&rows=10")
	fmt.Println(jsonStr)
	var tempData = make(map[string]interface{})
	json.Unmarshal([]byte(jsonStr), &tempData)
	fmt.Println(tempData)
	dataSlice := tempData["showapi_res_body"].(map[string]interface{})["data"].([]interface{})
	fmt.Printf("type=%T, value=%v", dataSlice, dataSlice)
	idiomsMap := make(map[string]Idiom)
	for i := range dataSlice {
		title := dataSlice[i].(map[string]interface{})["title"].(string)
		idiom := Idiom{Title: title}
		idiomsMap[title] = idiom
	}
	fmt.Println(idiomsMap)
}

func GetJson(url string) (jsonStr string, err error) {
	// 获得数据
	// 获得网络数据
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http请求失败，err=", err)
		return
	}
	// 延时关闭IO资源
	defer resp.Body.Close()
	// resp.Body实现了Reader接口，对其进行数据读入
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取网络失败，err=", err)
		return
	}
	// 将网络数据转化为字符串输出
	jsonStr = string(bytes)
	return
}
