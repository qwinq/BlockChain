package main

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockChainIterator struct {
	db *bolt.DB
	currentHashPointer []byte
}

func (bc *BlockChain)NewBcIterator() *BlockChainIterator {
	return &BlockChainIterator{bc.db,bc.tail}
}
func (it *BlockChainIterator)Next() (block *Block) {
	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(bcBucket))
		if bucket==nil{log.Panic("db.View : bucket不存在")}
		block = Deserialize(bucket.Get(it.currentHashPointer))

		it.currentHashPointer=block.PrevHash
		return nil
	})
	return block
}