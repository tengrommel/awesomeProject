package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

const reImg = `<img[\s\S]+?src="(http[\s\S]+?)"`

// 获得页面上所有的图片连接
func spiderImage(url string) []string {
	html := GetHtml(url)
	re := regexp.MustCompile(reImg)
	rets := re.FindAllStringSubmatch(html, -1)
	fmt.Println("捕获图片张数：", len(rets))
	imgUrls := make([]string, 0)
	for _, ret := range rets {
		imgURl := ret[1]
		imgUrls = append(imgUrls, imgURl)
	}
	return imgUrls
}

func GetHtml(url string) string {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}

func DownloadImg(url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err := ioutil.WriteFile(fmt.Sprintf("d/%v", strconv.Itoa(int(time.Now().UnixNano()))+".jpg"), imgBytes, 0644)
	if err != nil {
		fmt.Println("下载成功！")
	} else {
		fmt.Println("下载失败：", err.Error())
	}
}

func main() {
	imgUrls := spiderImage("http://www.163.com")
	for _, imgUrl := range imgUrls {
		fmt.Println(imgUrl)
		DownloadImg(imgUrl)
	}
}
