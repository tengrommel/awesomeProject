package main

import (
	"awesomeProject/m/context/ex2/log"
	"flag"
	"fmt"
	stdLog "log"
	"net/http"
	"time"
)

func main() {
	flag.Parse()
	http.HandleFunc("/", log.Decorate(handler))
	stdLog.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println(ctx, "handler started")
	defer log.Println(ctx, "handler ended")
	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println(ctx, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
