package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(" Replace 函数的用法")
	fmt.Println(strings.Replace("oink/oink/oink", "oink/oink", "ky", 1))
	h := md5.New()
	nowStr := strconv.Itoa(time.Now().Nanosecond())
	randomStr := strconv.Itoa(rand.Intn(10000))
	h.Write([]byte(nowStr + randomStr))
	fmt.Println(hex.EncodeToString(h.Sum(nil))[:12])

	str := "/a"
	nameList := strings.Split(str, "/")
	for i := range nameList {
		fmt.Println(nameList[i] == "", i)
	}
}
