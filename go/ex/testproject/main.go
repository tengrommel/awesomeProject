package main

import (
	"github.com/gorilla/mux"
	"github.com/labstack/gommon/log"
	"net/http"
)

func Add(value1 int, value2 int) int {
	return value1 + value2
}

func Subtract(value1 int, value2 int) int {
	return value1 - value2
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":12345", router))
}
