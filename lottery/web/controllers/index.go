package controllers

import (
	"../../models"
	"../../services"
	"github.com/kataras/iris"
)

type IndexController struct {
	Ctx         iris.Context
	ServiceGift services.GiftServer
}

func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type", "text/html")
	return "welcome to Go 抽奖"
}

func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{})
	rs["code"] = 0
	rs["msg"] = ""
	dataList := c.ServiceGift.GetAll()
	list := make([]models.Gift, 0)
	for _, data := range dataList {
		if data.SysStatus == 0 {
			list = append(list, data)
		}
	}
	rs["gifts"] = list
	return rs
}
