package main

import (
	"io"
	"log"
	"net/http"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// A very simple health check
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

// 文件服务器
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/health-check", HealthCheckHandler)
	log.Fatal(http.ListenAndServe(":9999", nil))
}
