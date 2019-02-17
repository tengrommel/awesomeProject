package main

import (
	"awesomeProject/notice/go_web/models"
	"awesomeProject/notice/go_web/routes"
	"awesomeProject/notice/go_web/utils"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	models.TestConnection()
	fmt.Printf("Listening Port %s\n", PORT)
	utils.LoadTemplates("views/*.html")
	r := routes.NewRouter()
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(PORT, nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
