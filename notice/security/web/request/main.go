package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}
	log.Fatal(server.ListenAndServe())
}
