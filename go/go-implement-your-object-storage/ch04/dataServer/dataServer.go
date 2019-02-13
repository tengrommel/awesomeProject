package main

import (
	"./heartbeat"
	"./locate"
	"log"
	"net/http"
	"os"
)

func main() {
	locate.CollectObjects()
	go locate.StartLocate()
	go heartbeat.StartHeartbeat()
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
