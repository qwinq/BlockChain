package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"log"
	"time"
)
//TODO 升级版
// 1. 补充区块字段
// 2. 更新计算哈希函数
// 3. 优化代码

// 1. 定义结构
type Block struct {
	Version uint64 		// 1. 版本号
	PrevHash []byte 	// 2. 前区块哈希
	MerkleRoot []byte	// 3. Merkle根
	TimeStamp uint64	// 4. 时间戳
	Difficulty uint64	// 5. 难度值
	Nonce uint64		// 6. 随机数(挖矿要找的数据)

	Hash []byte 		// a. 当前区块哈希(真币中无此项)
	Data []byte 		// b. 数据
}
// 实现辅助函数:uint64转[]byte
func Uint64ToByte(num uint64) []byte {
	buffer:=bytes.Buffer{}
	err := binary.Write(&buffer, binary.BigEndian, num)
	if err!=nil{
		log.Panic(err)
	}
	return buffer.Bytes()
}

// 2. 创建区块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := Block{
		Version:    0,
		PrevHash:   prevBlockHash,
		MerkleRoot: nil,
		TimeStamp:  uint64(time.Now().Unix()),
		Difficulty: 0, //填入无效值
		Nonce:      0, //同上
		Hash:       nil,
		Data:       []byte(data),
	}
	block.SetHash() // 后续添加(已完成)
	return &block
}
// 3. 生成哈希
func (block *Block)SetHash()  {
	//TO DO
	// 1. 拼装数据 (前区块hash+数据)
	/*
	var blockInfo []byte
	blockInfo=append(blockInfo,block.PrevHash...)
	blockInfo=append(blockInfo,Uint64ToByte(block.Version)...)
	blockInfo=append(blockInfo,block.MerkleRoot...)
	blockInfo=append(blockInfo,Uint64ToByte(block.TimeStamp)...)
	blockInfo=append(blockInfo,Uint64ToByte(block.Difficulty)...)
	blockInfo=append(blockInfo,Uint64ToByte(block.Nonce)...)
	blockInfo=append(blockInfo,block.Data...)
	*/
	// bytes.Join优化
	tmp:=[][]byte{
		block.PrevHash,
		Uint64ToByte(block.Version),
		block.MerkleRoot,
		Uint64ToByte(block.TimeStamp),
		Uint64ToByte(block.Difficulty),
		Uint64ToByte(block.Nonce),
		block.Data,
	}
	blockInfo := bytes.Join(tmp, []byte(""))
	// 2. sha256
	// func Sum256(data []byte) [Size]byte {
	hash:=sha256.Sum256(blockInfo)
	block.Hash=hash[:]
}
