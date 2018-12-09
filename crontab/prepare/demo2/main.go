package main

import (
	"fmt"
	"os/exec"
)

var (
	cmd    *exec.Cmd
	output []byte
	err    error
)

func main() {
	cmd = exec.Command("/bin/bash", "-c", "sleep 5;ls -l")
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(output))
}
