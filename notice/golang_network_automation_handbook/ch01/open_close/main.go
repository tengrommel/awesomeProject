package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write([]byte("append data\n"))
}
