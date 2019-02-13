package main

import (
	"./heartbeat"
	"./versions"
	"awesomeProject/go/go-implement-your-object-storage/ch04/apiServer/locate"
	"awesomeProject/go/go-implement-your-object-storage/ch04/apiServer/objects"
	"log"
	"net/http"
	"os"
)

func main() {
	go heartbeat.ListenHeartbeat()
	http.HandleFunc("/objects/", objects.Handler)
	http.HandleFunc("/locate/", locate.Handler)
	http.HandleFunc("/versions/", versions.Handler)
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil))
}
