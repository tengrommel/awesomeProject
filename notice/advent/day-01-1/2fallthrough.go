package main

import "fmt"

func main() {
	for {
		fmt.Println("请输入你的评级：A/B/C")
		var level string
		fmt.Scan(&level)
		switch level {
		case "A":
			fmt.Println("恭喜您获得我厂生产的女朋友一枚")
			fallthrough
		case "B":
			fmt.Println("恭喜您获得书包一只")
			fallthrough
		case "C":
			fmt.Println("恭喜您获得水杯一个")
		case "fuckOff":
			return
		}
	}

}
