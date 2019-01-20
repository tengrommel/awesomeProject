package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"time"
)

// 红包列表
var packageList map[uint32][]uint

type lotteryController struct {
	Ctx iris.Context
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

// 返回全部红包地址
// http://localhost:8080/
func (c *lotteryController) Get() map[uint32][2]int {
	rs := make(map[uint32][2]int)
	for id, list := range packageList {
		var money int
		for _, v := range list {
			money += int(v)
		}
		rs[id] = [2]int{len(list), money}
	}
	return rs
}

func (c *lotteryController) GetSet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	money, errMoney := c.Ctx.URLParamFloat64("money")
	num, errNum := c.Ctx.URLParamInt("num")
	if errUid != nil || errMoney != nil || errNum != nil {
		return fmt.Sprintf("errUid=%d, errMoney=%d, errNum=%d\n", errUid, errMoney, errNum)
	}
	moneyTotal := int(money * 100)
	if uid < 1 || moneyTotal < num || num < 1 {
		return fmt.Sprintf("参数数值异常，uid=%d, money=%d, num=%d\n", uid, money, num)
	}
	// 金额分配算法
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	rMax := 0.55 // 随机分配最大值
	list := make([]uint, num)
	leftMoney := moneyTotal
	leftNum := num
	// 大循环开始，分配金额给到每一个红包
	for leftNum > 0 {
		if leftNum == 1 {
			// 最后一个红包，剩余的全部金额都给它
			list[num-1] = uint(leftMoney)
			break
		}
		if leftMoney == leftNum {
			for i := num - leftNum; i < num; i++ {
				list[i] = 1
			}
			break
		}
		rMoney := int(float64(leftMoney-leftNum) * rMax)
		m := r.Intn(rMoney)
		if m < 1 {
			m = 1
		}
		list[num-leftNum] = uint(m)
		leftMoney -= m
		leftNum--
	}
	// 红包的唯一ID
	id := r.Uint32()
	packageList[id] = list
	// 返回红包的URL
	return fmt.Sprintf("/get?id=%d&uid=%d&num=%d", id, uid, num)
}

func (c *lotteryController) GetGet() string {
	uid, errUid := c.Ctx.URLParamInt("uid")
	id, errId := c.Ctx.URLParamInt("id")
	if errUid != nil || errId != nil {
		return fmt.Sprintf("")
	}
	if uid < 1 || id < 1 {
		return fmt.Sprintf("")
	}
	list, ok := packageList[uint32(id)]
	if !ok || len(list) < 1 {
		return fmt.Sprintf("红包不存在，id=%d\n", id)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	i := r.Intn(len(list))
	money := list[i]
	// 更新红包列表下的信息
	if len(list) > 1 {
		if i == len(list)-1 {
			packageList[uint32(id)] = list[:i]
		} else if i == 0 {
			packageList[uint32(id)] = list[1:]
		} else {
			packageList[uint32(id)] = append(list[:i], list[i+1:]...)
		}
	} else {
		delete(packageList, uint32(uid))
	}
	return fmt.Sprintln("恭喜你抢到红包：", money)
}
