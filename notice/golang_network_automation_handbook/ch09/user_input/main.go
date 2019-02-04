package main

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"syscall"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username:")
	user, _ := reader.ReadString('\n')
	fmt.Print("Enter password:")
	pass, _ := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Printf("\n\nUsername: %sPassword:%s\n", user, string(pass))
}
