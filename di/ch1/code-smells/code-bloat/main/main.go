package main

import (
	"fmt"
	"github.com/tengrommel/awesomeProject/di/ch1/code-smells/code-bloat"
)

func main() {
	code_bloat.MakeListAgain()
	fmt.Println(code_bloat.StringList)
}
