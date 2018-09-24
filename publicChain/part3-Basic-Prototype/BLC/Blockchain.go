package BLC

type Blockchain struct {
	Blocks []*Block // 存储有序的区块
}

// 1、创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock() *Blockchain {
	genesisBlock := CreateGenesisBlock("Genesis Data......")
	return &Blockchain{[]*Block{genesisBlock}}
}
