package main

import (
	"bufio"
	"os"
)

func main() {
	lines := []string{
		"RTR1 1.1.1.1",
		"RTR1 2.2.2.2",
		"RTR1 3.3.3.3"}
	file, _ := os.OpenFile("./file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, _ = writer.WriteString(line + "\n")
	}
	writer.Flush()
}
