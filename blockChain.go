package main

import (
	"github.com/boltdb/bolt"
	"log"
)

// 4. 引入区块链
// V3
// 1. BlockChain结构重写 使用数据库代替切片
// 1.1. 数据结构设计:存储结构
// 		a. key=block.hash value=block.Serialize
// 		b. 记录末位区块hash: lastHash->block.Hash
// 2. NewBlockChain函数重写 由数组操作改写成==对数据库操作==，创建数据库
// 3. AddBlock函数重写 对数据库的读取和写入
type BlockChain struct {
	// 定义区块链数组
	//blocks []*Block
	db *bolt.DB // 数据库代替数组
	tail []byte // lastBLock.Hash
}
const (
	bcDB     = "BlockChain.db"
	bcBucket ="bcBucket"
	lastHashKey="LastHashKey"
)
// 5. 定义区块链
func NewBlockChain() *BlockChain {
	//genesisBlock := GenesisBlock()
	//return &BlockChain{blocks:[]*Block{genesisBlock}}
	// 切片操作->数据库操作,创建数据库
	db, err := bolt.Open(bcDB, 0600, nil)
	defer db.Close()
	if err != nil {
		log.Panic("bolt.OPen:", err)
	}

	var lastHash []byte

	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bcBucket))
		if bucket == nil {
			// 无则创建
			bucket, err = tx.CreateBucket([]byte(bcBucket))
			if err != nil {log.Panic("tx.Create: ", err)}
			genesisBlock := GenesisBlock()
			bucket.Put(genesisBlock.Hash,genesisBlock.Serialize())
			bucket.Put([]byte(lastHashKey),genesisBlock.Hash)
			lastHash=genesisBlock.Hash
			//fmt.Printf("%v",Deserialize(bucket.Get(lastHash)))
		}else{
			lastHash = bucket.Get([]byte(lastHashKey))
		}
		return nil
	})

	return &BlockChain{db, lastHash,}
}

func OpenDB() *bolt.DB {
	db, err := bolt.Open(bcDB, 0600, nil)
	if err != nil {
		log.Panic("bolt.OPen:", err)
	}
	return db
}
func ViewFromDB(db *bolt.DB,tail []byte)(bytes []byte) {
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bcBucket))
		if bucket==nil{log.Panic("open bucket err")}
		bytes = bucket.Get(tail)
		return nil
	})
	return
}
func Write2DB(db *bolt.DB, bc BlockChain, block *Block) {
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bcBucket))
		if bucket == nil {
			// 无则创建
			bucket, _ = tx.CreateBucket([]byte(bcBucket))
			//if err != nil {log.Panic("tx.Create: ", err)}
		}
		bc.db = db
		bc.tail = block.Hash //记录末位区块hash: lastHash->block.Hash
		//写入创世区块 a. key=block.hash value=block.Serialize
		bucket.Put(block.Hash, block.Data)
		bucket.Put([]byte("lastHash"), bc.tail)
		return nil
	})
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
	//lastBlock:=bc.blocks[len(bc.blocks)-1]
	//prevHash:=lastBlock.Hash
	//block:=NewBlock(data,prevHash)
	// b. 添加到区块链数组中
	//bc.blocks= append(bc.blocks, block)

	// a. 创建新的区块
	// 获取区块链的末尾区块
	block:=NewBlock(data,bc.tail)
	// b. 添加到区块链数组中
	Write2DB(bc.db,*bc,block)
}

