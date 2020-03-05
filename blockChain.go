package main

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
