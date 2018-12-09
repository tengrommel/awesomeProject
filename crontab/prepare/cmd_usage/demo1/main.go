package main

import (
	"fmt"
	"os/exec"
)

var (
	cmd *exec.Cmd
	err error
)

func main() {
	cmd = exec.Command("/bin/zsh", "-c", "echo 1")
	err = cmd.Run()
	fmt.Println(err)
}
