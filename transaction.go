package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

// 1. 定义交易结构 设置交易ID
// 2. 构建交易
// 3. 创建挖矿交易
// 4. 根据交易调整程序
type transaction struct {
	TXID []byte			//交易ID
	TXInputs []TXInput	//交易输入切片
	TXOutputs []TXOutput	//交易输入切片
}
type TXInput struct {
	// 引用的交易ID
	Index int64
	// 引用的output索引
	TXid []byte
	// 解锁脚本(用地址模拟)
	Sig string
}
type TXOutput struct {
	// 转账金额
	Value float64
	// 锁定脚本(地址模拟)
	PubKeyHash string
}
// 设置交易ID
func (tx transaction)SetHash() {
	buffer:=bytes.Buffer{}
	if err := gob.NewEncoder(&buffer).Encode(tx);err!=nil{
		log.Panic(err)
	}
	data:=buffer.Bytes()
	hash:=sha256.Sum256(data)
	tx.TXID=hash[:]

}
