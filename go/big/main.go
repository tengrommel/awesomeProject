package main

import (
	"bufio"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"os"
)

func main() {
	//contentBytes, err := ioutil.ReadFile("kai.txt")
	//if err != nil {
	//	fmt.Println("读入失败", err)
	//}
	//contentStr := string(contentBytes)
	//// 逐行打印
	//lineStrs := strings.Split(contentStr, "\n")
	//for _, lineStr := range lineStrs {
	//	fmt.Println(lineStr)
	//	newStr, _ := convertEncoding(lineStr, "GBK", "UTF-8")
	//	fmt.Println(newStr)
	//}
	file, _ := os.Open("file.txt")
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		lineBytes, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		gbkStr := string(lineBytes)
		utfStr, _ := convertEncoding(gbkStr, "GBK", "UTF-8")
		fmt.Println(utfStr)
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
