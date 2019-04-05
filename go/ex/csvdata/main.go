package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// Person -
type Person struct {
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Address   *Address `json:"address"`
}

// Address -
type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

func main() {
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(csvFile)
	var people []Person
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		people = append(people, Person{
			FirstName: line[0],
			LastName:  line[1],
			Address: &Address{
				City:  line[2],
				State: line[3],
			},
		})
	}
	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))
}
