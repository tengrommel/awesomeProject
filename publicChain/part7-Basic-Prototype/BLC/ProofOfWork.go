package BLC

import "math/big"

// 0000 0000 0000 0000 1001 0002 0000 ... 0001

// 256位Hash里面前面至少要有16个零
const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	// 1、将Block的属性拼接成字节数组
	// 2、生成hash
	// 3、判断hash有效性，如果满足条件，跳出循环
	return nil, 0
}

// 创建新的工作量证明对象
func NewProofOfWork(block *Block) *ProofOfWork {
	// 1、big.Int对象 1
	// 2、
	// 0000 0001
	// 8 - 2 = 6
	// 0100 0000 64
	// 0010 0000
	// 0000 0000 0000 0001 0000 0000 0000 0000 ... 000
	// 1、创建一个初始值位1的target
	target := big.NewInt(1)
	// 2、左移256 - targetBit
	target = target.Lsh(target, 256-targetBit)

	return &ProofOfWork{block, target}
}
