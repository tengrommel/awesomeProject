package main

import (
	"fmt"
	"github.com/axgle/mahonia"
	"io/ioutil"
	"strings"
)

func main() {
	contentBytes, err := ioutil.ReadFile("kai.txt")
	if err != nil {
		fmt.Println("读入失败", err)
	}
	contentStr := string(contentBytes)
	// 逐行打印
	lineStrs := strings.Split(contentStr, "\n")
	for _, lineStr := range lineStrs {
		fmt.Println(lineStr)
		newStr, _ := convertEncoding(lineStr, "GBK", "UTF-8")
		fmt.Println(newStr)
	}
}

func convertEncoding(srcStr string, srcEncoding string, dstEncoding string) (string, error) {
	srcDecoder := mahonia.NewDecoder(srcEncoding)
	dstDecoder := mahonia.NewDecoder(dstEncoding)
	utfStr := srcDecoder.ConvertString(srcStr)
	_, dstBytes, err := dstDecoder.Translate([]byte(utfStr), true)
	if err != nil {
		return "", err
	}
	return string(dstBytes), nil
}
