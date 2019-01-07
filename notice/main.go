package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.planetmeican.com/planet/go-tool/httputil"
)

type JSONResponse struct {
	ResultCode        string      `json:"resultCode"`
	ResultDescription string      `json:"resultDescription"`
	Data              interface{} `json:"data"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/ok", func(context *gin.Context) {
		fmt.Println(context.Request)
		f := struct {
			ClientUUID  string `json:"clientUUID"`
			ReportUUID  string `json:"reportUUID"`
			DownloadUrl string `json:"downloadUrl"`
		}{}
		if err := context.ShouldBindJSON(&f); err != nil {
			context.JSON(200, err.Error())
			return
		}
		fmt.Println(f)
		context.JSON(200, JSONResponse{
			ResultCode: httputil.OK,
		})
	})
	err := r.Run(":9999") // listen and serve on 0.0.0.0:8080
	panic(err)
}
