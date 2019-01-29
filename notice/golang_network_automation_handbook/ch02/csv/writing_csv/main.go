package main

import (
	"encoding/csv"
	"os"
)

func main() {
	rows := [][]string{
		{"HOST", "IP ADDR", "LOCATION"},
		{"RTR1", "1.1.1.1", "Seattle, WA"},
		{"RTR2", "2.2.2.2", "Nevada,NV"}}
	file, _ := os.Create("file.csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, row := range rows {
		_ = writer.Write(row)
	}
	writer.Flush()
}
