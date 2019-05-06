package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type DataTO struct {
	Time  time.Time `json:"time"`
	Key   string    `json:"key"`
	Name  string    `json:"name"`
	Value int       `json:"value"`
	Son   *[]DataTO `json:"son"`
}

type Hello struct{}

func (h *Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("this is next door!")
	itemList := make([]DataTO, 0)
	for i := 0; i < 10; i++ {
		var key string
		if i%2 == 0 {
			key = "ok"
		} else {
			key = "error"
		}
		itemList = append(itemList, DataTO{
			Time:  time.Now(),
			Name:  strconv.FormatInt(int64(i), 10),
			Value: i,
			Key:   key,
		})
	}
	dataTO := DataTO{
		Key:   "key",
		Name:  "teng",
		Value: 13,
		Son:   &itemList,
	}
	res := Response{
		200,
		"success",
		dataTO,
	}
	result, _ := json.Marshal(res)
	w.Write(result)
}

type Response struct {
	Errno int         `json:"errno"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func main() {
	http.HandleFunc("/web", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "fmt print is ok!")
	})
	h := Hello{}
	http.Handle("/hello", &h)
	log.Fatal(http.ListenAndServe(":9080", nil))
}
