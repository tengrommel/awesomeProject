package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Since the goal is to read each line of text,
bufio.ScanLines is leveraged as an input to the bufio.ScanLines()
method and then the Scanner advances to each new line using the bufio.Scanner.Scan() method and then
the Scanner advances to each new line using the bufio.Scanner.Scan() method.
*/

func main() {
	file, _ := os.Open("./file.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
