package main

import "fmt"

//TODO 简单版
// 1. 定义结构  1. 前区块哈希 2. 当前区块哈希 3. 数据
// 2. 创建区块
// 3. 生成哈希
// 4. 引入区块链
// 5 .添加区块
// 6. 重构代码
// 1. 定义结构
type Block struct {
	// 1. 前区块哈希 2. 当前区块哈希 3. 数据
	PrevHash, Hash, Data []byte
}

// 2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		PrevHash: prevBlockHash,
		Hash:     []byte{}, //预制为空,后续添加 // TODO
		Data:     []byte(data),
	}
	return &block
}

//TODO 升级版
// 1. 补充区块字段
// 2. 更新计算哈希函数
// 3. 优化代码
func main() {
	block := NewBlock("first bitCoin", []byte{})
	fmt.Printf("前区块哈希值:%x\n", block.PrevHash)
	fmt.Printf("当前区块哈希值:%x\n", block.Hash)
	fmt.Printf("数据:%s\n", string(block.Data)) //right now
}
