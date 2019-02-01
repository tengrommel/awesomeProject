package main

import (
	"awesomeProject/go/go-implement-your-object-storage/ch03/apiServer/heartbeat"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
