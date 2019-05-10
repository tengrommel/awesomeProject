package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Developer bool `json:"is_developer"`
}

func main() {
	fmt.Println("Hello World")
	author := Author{Name: "Elliot Forbes", Age: 25, Developer: true}
	fmt.Println(author)
	book := Book{Title: "Learning Concurrency in Python", Author: "Elliot Forbes"}
	fmt.Printf("%+v\n", book)
	byteArray, err := json.MarshalIndent(book, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteArray))

}
