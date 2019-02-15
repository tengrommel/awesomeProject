package main

import (
	"fmt"
	"go.planetmeican.com/titan/log"
	"net/http"
	"reflect"
	"runtime"
)

type HelloHandler struct{}

func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func midLog(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h.ServeHTTP(w, r)
	})
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Info("protect")
		h.ServeHTTP(writer, request)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9999",
	}
	hello := HelloHandler{}
	http.Handle("/hello", protect(midLog(hello)))
	log.Fatal(server.ListenAndServeTLS("/Users/tengzhou/go/src/awesomeProject/cert.pem",
		"/Users/tengzhou/go/src/awesomeProject/key.pem"))
}
