package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "r.txt", "w.txt")
	data, _ := cmd.CombinedOutput()
	fmt.Printf("%s", data)
}
