package main

import (
	"awesomeProject/tt/business_demo/poms/ctrl"
	"fmt"
	"log"
	"net/http"
)

func main() {
	ctrl.Setup()
	go http.ListenAndServe(":3000", new())
	log.Println("Server started, press <ENTER> to exit")
	fmt.Scanln()
}
