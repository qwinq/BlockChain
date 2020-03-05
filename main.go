package main

import (
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
