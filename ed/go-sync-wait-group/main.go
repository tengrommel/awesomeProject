package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"https://google.com",
	"https://tutorialedge.net",
	"https://twitter.com",
}

func fetchStatus(writer http.ResponseWriter, request *http.Request) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(writer, "%+v\n", err)
			}
			fmt.Fprintf(writer, "%+v\n", resp)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func main() {
	fmt.Println("Go WaitGroup Tutorial")
	http.HandleFunc("/", fetchStatus)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
