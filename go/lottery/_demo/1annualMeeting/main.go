package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"strings"
	"time"
)

var userList []string

type lotteryControl struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryControl{})
	return app
}

func main() {
	app := newApp()
	userList = make([]string, 0)
	log.Fatal(app.Run(iris.Addr(":8080")))
}

func (c *lotteryControl) Get() string {
	count := len(userList)
	return fmt.Sprintf("当前总共参与抽奖的用户数： %d\n", count)
}

func (c *lotteryControl) PostImport() string {
	strUsers := c.Ctx.FormValue("users")
	users := strings.Split(strUsers, ",")
	count1 := len(userList)
	for _, u := range users {
		u = strings.TrimSpace(u)
		if len(u) > 0 {
			userList = append(userList, u)
		}
	}
	count2 := len(userList)
	return fmt.Sprintf("当前总共参与的用户数: %d, 成功导入的用户数：%d\n", count2, (count2 - count1))
}

func (c *lotteryControl) GetLucky() string {
	count := len(userList)
	if count > 1 {
		seed := time.Now().UnixNano()
		index := rand.New(rand.NewSource(seed)).Int31n(int32(count))
		user := userList[index]
		userList = append(userList[0:index], userList[index+1:]...)
		return fmt.Sprintf("当前中奖用户：%s，剩余用户数：%d\n", user, count-1)
	} else if count == 1 {
		user := userList[0]
		return fmt.Sprintf("当前中奖用户： %s, 剩余用户数：%d\n", user, count-1)
	} else {
		return fmt.Sprintf("已经没有参与用户，清先通过 /import 导入用户")
	}
}
