package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := exec.Command("cat", "r.txt", "w.txt")
	stdoutpipe, _ := cmd.StdoutPipe()
	stderrpipe, _ := cmd.StderrPipe()
	_ = cmd.Start()
	stdout, _ := ioutil.ReadAll(stdoutpipe)
	stderr, _ := ioutil.ReadAll(stderrpipe)
	fmt.Printf("stdout: %sstderr: %s", stdout, stderr)
	_ = cmd.Wait()
}
