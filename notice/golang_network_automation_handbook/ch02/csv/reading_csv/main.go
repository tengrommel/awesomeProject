package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

/*
The csv package includes a NewReader() function that returns a Reader object for processing CSV records.
A csv.Reader converts all \r\n sequences in its input to just \n, including multiline field values.
*/

type Row struct {
	ColA, ColB, ColC string
}

func main() {
	file, _ := os.Open("file.csv")
	reader := csv.NewReader(file)
	rows := []Row{}
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		rows = append(rows, Row{
			ColA: row[0],
			ColB: row[1],
			ColC: row[2],
		})
	}
	for _, row := range rows {
		fmt.Printf("%s--%s--%s\n", row.ColA, row.ColB, row.ColC)
	}
}
