package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/gocarina/gocsv"
	"github.com/jszwec/csvutil"
	"io"
	"log"
	"strings"
)

type Client struct { // Our example struct, you can use "-" to ignore a field
	Id      string `csv:"client_id"`
	Name    string `csv:"client_name"`
	Age     string `csv:"client_age"`
	NotUsed string `csv:"-"`
}

func main() {
	var buf bytes.Buffer
	clients := []*Client{}
	clients = append(clients, &Client{Id: "12", Name: "John", Age: "21"}) // Add clients
	clients = append(clients, &Client{Id: "13", Name: "Fred"})
	clients = append(clients, &Client{Id: "14", Name: "James", Age: "32"})
	clients = append(clients, &Client{Id: "15", Name: "Danny"})
	err := gocsv.Marshal(&clients, &buf) // Get all clients as CSV string
	if err != nil {
		panic(err)
	}
	fmt.Println(string(buf.Bytes())) // Display all clients as CSV string

	type User struct {
		ID   int
		Name string
		Age  int `csv:",omitempty"`
		City string
	}

	csvReader := csv.NewReader(strings.NewReader(`
1,John,27,la
2,Bob,,ny`))

	// in real application this should be done once in init function.
	userHeader, err := csvutil.Header(User{}, "csv")
	if err != nil {
		log.Fatal(err)
	}

	dec, err := csvutil.NewDecoder(csvReader, userHeader...)
	if err != nil {
		log.Fatal(err)
	}

	var users []User
	for {
		var u User
		if err := dec.Decode(&u); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	fmt.Printf("%+v", users)
}
