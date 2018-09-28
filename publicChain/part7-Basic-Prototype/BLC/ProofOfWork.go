package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

// 0000 0000 0000 0000 1001 0002 0000 ... 0001

// 256位Hash里面前面至少要有16个零
const targetBit = 16

type ProofOfWork struct {
	Block  *Block
	target *big.Int
}

func (pow *ProofOfWork) prepareData(nonce int64) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevBlockHash,
			pow.Block.Data,
			IntToHex(pow.Block.Timestamp),
			IntToHex(int64(targetBit)),
			IntToHex(int64(nonce)),
			IntToHex(int64(pow.Block.Height)),
		},
		[]byte{})
	return data
}

func (proofOfWork *ProofOfWork) Run() ([]byte, int64) {
	// 1、将Block的属性拼接成字节数组
	// 2、生成hash
	// 3、判断hash有效性，如果满足条件，跳出循环
	nonce := int64(0)
	var hashInt big.Int // 存储我们新生成的hash
	var hash [32]byte
	for {
		// 准备数据
		dataBytes := proofOfWork.prepareData(nonce)
		// 生成hash
		hash = sha256.Sum256(dataBytes)
		fmt.Printf("\r%x\n", hash)
		// 将hash存储到hashInt
		hashInt.SetBytes(hash[:])
		// 判断hashInt是否小于Block里面的target
		if proofOfWork.target.Cmp(&hashInt) == 1 {
			break
		}
		nonce = nonce + 1
	}
	return hash[:], nonce
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
