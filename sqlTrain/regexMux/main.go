package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func main() {
	http.HandleFunc("/", router)
	//http.HandleFunc("/subpath", hdlSub)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

var numbericPath, _ = regexp.Compile("[0-9]")
var numbericSuffic, _ = regexp.Compile("[0-9]$")
var numbericPrefix, _ = regexp.Compile("^/[0-9]")

func router(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	switch {
	case numbericPath.MatchString(r.URL.Path):
		hdlNum(w, r)
	case numbericPrefix.MatchString(r.URL.Path):
		hdlPrefix(w, r)
	case numbericSuffic.MatchString(r.URL.Path):
		hdlNumSuf(w, r)
	default:
		hdlMain(w, r)
	}
}

func hdlPrefix(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on prefix handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlNum(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on numeric path handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlNumSuf(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on numeric suffix path handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlMain(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on default handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}

func hdlSub(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received on sub handler ...")
	content, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(content))
}
