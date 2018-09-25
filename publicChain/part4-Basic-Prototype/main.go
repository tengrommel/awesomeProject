package main

import (
	"awesomeProject/publicChain/part4-Basic-Prototype/BLC"
	"fmt"
)

func main() {
	// 创世区块
	blockchain := BLC.CreateBlockchainWithGenesisBlock()
	fmt.Println(blockchain.Blocks[0])
	// 新区块
	blockchain.AddBlockToBlockchain("Send 100RMB To zhangqiang",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 200RMB To zhouteng",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 300RMB To teng",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	blockchain.AddBlockToBlockchain("Send 50RMB To ten",
		blockchain.Blocks[len(blockchain.Blocks)-1].Height+1,
		blockchain.Blocks[len(blockchain.Blocks)-1].Hash)
	fmt.Println(blockchain)
	fmt.Println(blockchain.Blocks)
}
