package main

import (
	"awesomeProject/publicChain/part2-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	genesisBlock := BLC.CreateGenesisBlock("Genesis Block...")
	fmt.Println(genesisBlock)
}
