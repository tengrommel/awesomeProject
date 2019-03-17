package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strconv"
	"sync"
	"time"
)

var (
	chSem      = make(chan int, 100)
	downloadWG sync.WaitGroup
	randomMT   sync.Mutex
)

const reImg = `<img[\s\S]+?src="(http[\s\S]+?)"`

// 生成[start, end)随机数
func GetRandomInt(start, end int) int {
	randomMT.Lock()
	defer randomMT.Unlock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := start + r.Intn(end-start)
	return ret
}

/*
生成时间戳_随机数文件名
*/
func GetRandomName() string {
	timestamp := strconv.Itoa(int(time.Now().UnixNano()))
	randomNum := strconv.Itoa(GetRandomInt(1000, 10000))
	return timestamp + randomNum
}

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
	fmt.Println("DownloadImg...")
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	imgBytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(fmt.Sprintf("/home/teng/go/src/awesomeProject/notice/advent/%v", GetRandomName()+".jpg"), imgBytes, 0644)
	if err == nil {
		fmt.Println("下载成功！")
	} else {
		fmt.Println("下载失败：")
	}
}

func DownloadImgAsync(url string) {
	downloadWG.Add(1)
	go func() {
		chSem <- 123
		DownloadImg(url)
		<-chSem
		downloadWG.Done()
	}()
	downloadWG.Wait()
}

func main() {
	imgUrls := spiderImage("http://www.163.com")
	for _, imgUrl := range imgUrls {
		fmt.Println(imgUrl)
		DownloadImgAsync(imgUrl)
	}
}
