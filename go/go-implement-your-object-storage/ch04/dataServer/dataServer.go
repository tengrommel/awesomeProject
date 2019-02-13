package main

import (
	"./heartbeat"
	"./locate"
	"./objects"
	"./temp"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go locate.StartLocate()
	go heartbeat.StartHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/temp/", temp.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
