package main

import "crypto/sha256"

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