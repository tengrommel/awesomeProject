package main

import (
	"./insertData"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	SERVER_PORT = ":8080"
)

/*
	end point for entering a/mu game results
	check if the request json doesn't have any errors and
	then insert the data to the database
*/

func insertDataEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	jsonBody, _ := insertData.InsertData(r)
	w.Write(jsonBody)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/insert", insertDataEndPoint).Methods("POST")
	http.ListenAndServe(SERVER_PORT, router)
}
