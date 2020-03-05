package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

//TODO
// 1. 定义工作量结构ProofOfWork
// a. block
// b. 目标值
// 2. POW函数 NewProofOfWork(参数)
// 3. hash函数(不断计算) Run()
// 4. 校验函数 IsValid()

// 1. 定义工作量结构ProofOfWork
type ProofOfWork struct {
	block *Block	// a. block
	target *big.Int	// b. 目标值 big.Int(比较,赋值等方法)
}
// 2. POW函数 NewProofOfWork(参数)
func NewProofOfWork(block *Block) *ProofOfWork{
	pow:=ProofOfWork{
		block:  block,
		//target: nil,
	}
	// 指定的难度值
	targetStr := "0000f00000000000000000000000000000000000000000000000000000000000"
	// 辅助变量 big.Int{}.SetString(str,16)
	tmpInt:=big.Int{}
	//tmpInt.SetString(targetStr, 16)
	//pow.target=&tmpInt
	if strPtr, ok := tmpInt.SetString(targetStr, 16);ok{
		pow.target=strPtr
	}
	return &pow
}
// 3. hash函数(不断计算) Run()
func (pow *ProofOfWork)Run() ([]byte,uint64) {
	// 1. 拼装数据(区块数据,随机数)
	// 2. hash运算
	// 3. 与pow.target比较 a.找到退出 b.没找到,随机数+1,继续找
	block := pow.block
	var (
		hash [32]byte
		nonce uint64
	)

	for {
		// 1. 拼装数据(区块数据,随机数)
		tmp:=[][]byte{
			block.PrevHash,
			Uint64ToByte(block.Version),
			block.MerkleRoot,
			Uint64ToByte(block.TimeStamp),
			Uint64ToByte(block.Difficulty),
			Uint64ToByte(nonce),
			block.Data,
		}
		blockInfo := bytes.Join(tmp, []byte(""))
		// 2. hash运算
		hash = sha256.Sum256(blockInfo)
		// 3. 与pow.target比较 a.找到退出 b.没找到,随机数+1,继续找
		tmpInt:=big.Int{}
		bigHash := tmpInt.SetBytes(hash[:]) // [32]byte转big.Int
		//   -1 if x <  y
		//    0 if x == y
		//   +1 if x >  y
		// func (x *Int) Cmp(y *Int) (r int) {
		if bigHash.Cmp(pow.target)==-1{
			// a.找到退出
			fmt.Printf("挖矿成功!hash : %x, nonce : %d\n",hash,nonce)
			break
		}else{
			// b.没找到,随机数+1,继续找
			nonce++
		}

	}
	return hash[:],nonce
}

