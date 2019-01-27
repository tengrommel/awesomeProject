package main

import (
	"awesomeProject/m/interfaces/retriever/mock"
	"awesomeProject/m/interfaces/retriever/real"
	"bufio"
	"fmt"
	"io"
	"time"
)

const url = "http:/www.imooc.com"

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

func download(r Retriever) string {
	return r.Get("http://www.imooc.com")
}

func post(poster Poster) {
	poster.Post(url,
		map[string]string{
			"name":   "ccmouse",
			"course": "golang",
		})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another faked imooc.com",
	})
	return s.Get(url)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is a fake imooc.com"}
	r = &retriever
	inspect(r)
	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	// Type assertion
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}
	//fmt.Println(download(r))
	fmt.Println("Try a session")
	fmt.Println(session(&retriever))
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf("%T %v\n", r, r)
	fmt.Print(" > Type switch:")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println(r)
}
