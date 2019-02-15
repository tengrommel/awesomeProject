package main

import (
	"log"
	"net/http"
)

// 文件服务器
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":9999", nil))
}
