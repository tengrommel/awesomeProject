package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func midLog(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", midLog(hello))
	log.Fatal(server.ListenAndServeTLS("/Users/tengzhou/go/src/awesomeProject/cert.pem",
		"/Users/tengzhou/go/src/awesomeProject/key.pem"))
}
