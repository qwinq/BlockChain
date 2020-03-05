package main

import (
	"crypto/sha256"
	"fmt"
)

//TODO 简单版
// 1. 定义结构  1. 前区块哈希 2. 当前区块哈希 3. 数据
// 2. 创建区块
// 3. 生成哈希
// 4. 引入区块链
// 5. 定义区块链
// 6. 添加区块
// 7. 重构代码

// 1. 定义结构
type Block struct {
	// 1. 前区块哈希 2. 当前区块哈希 3. 数据
	PrevHash, Hash, Data []byte
}

// 2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		//Hash:     []byte{}, //预制为空,后续添加 // TO DO
		Data:     []byte(data),
	}
	block.SetHash() // 后续添加(已完成)
	return &block
}
// 3. 生成哈希
func (block *Block)SetHash()  {
	//TO DO
	// 1. 拼装数据 (前区块hash+数据)
	blockInfo:=append(block.PrevHash,block.Data...)
	// 2. sha256
	// func Sum256(data []byte) [Size]byte {
	hash:=sha256.Sum256(blockInfo)
	block.Hash=hash[:]
}
// 4. 引入区块链
type BlockChain struct {
	// 定义区块链数组
	blocks []*Block
}
// 5. 定义区块链
func NewBlockChain() *BlockChain {
	genesisBlock := GenesisBlock()
	return &BlockChain{blocks:[]*Block{genesisBlock}}
}
// 5.1 定义创世区块
func GenesisBlock() *Block {
	return NewBlock("GenesisBlock!!!",[]byte{})
}
// 6. 添加区块
// a. 创建新的区块
// b. 添加到区块链数组中
func (bc *BlockChain)AddBlock(data string)  {
	// a. 创建新的区块
	// 获取区块链的末尾区块
	lastBlock:=bc.blocks[len(bc.blocks)-1]
	prevHash:=lastBlock.Hash
	block:=NewBlock(data,prevHash)
	// b. 添加到区块链数组中
	bc.blocks= append(bc.blocks, block)
}
//TODO 升级版
// 1. 补充区块字段
// 2. 更新计算哈希函数
// 3. 优化代码
func main() {
	bc := NewBlockChain()
	bc.AddBlock("Second Block")
	bc.AddBlock("Third Block")

	for i,block:=range bc.blocks{
		fmt.Printf("当前高度:%d\n", i)
		fmt.Printf("前区块哈希值:%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值:%x\n", block.Hash)
		fmt.Printf("数据:%s\n", string(block.Data))
	}


}
