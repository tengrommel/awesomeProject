package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Fatal(fmt.Fprintf(writer, "Hello, World"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
