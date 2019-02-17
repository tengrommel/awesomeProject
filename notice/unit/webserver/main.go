package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/double", doubleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func doubleHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue("v")
	if text == "" {
		http.Error(w, "missing value", http.StatusBadRequest)
		return
	}
	v, err := strconv.Atoi(text)
	if err != nil {
		http.Error(w, "not a number:"+text, http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, v*2)
}
