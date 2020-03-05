package main

import "math/big"

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
	targetStr := "0000100000000000000000000000000000000000000000000000000000000000"
	// 辅助变量 big.Int{}.SetString(str,16)
	tmpInt:=big.Int{}
	//tmpInt.SetString(targetStr, 16)
	//pow.target=&tmpInt
	if strPtr, ok := tmpInt.SetString(targetStr, 16);ok{
		pow.target=strPtr
	}

	return &pow
}


