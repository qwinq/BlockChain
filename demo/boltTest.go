package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main1() {
	// 1. 打开数据库
	db, err := bolt.Open("boltTest.db", 0600, nil)
	defer db.Close()
	if err!=nil{ log.Panic("打开bolt失败") }
	// 2. 更新数据库(无则建桶,有则写入)
	bucketName:=[]byte("b1")
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket==nil{
			// 2.1 无抽屉,创建抽屉
			bucket, err = tx.CreateBucket(bucketName)
			if err!=nil{log.Panic("创建bolt失败")}
		}
		// 2.2 有抽屉,写数据
		key:=[]byte("1")
		value:=[]byte("hello")
		bucket.Put(key,value)
		bucket.Put([]byte("2"),[]byte("world"))
		return nil
	})
	// 3. 读数据(无桶报错,有则读取)
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucketName)
		if bucket==nil{
			log.Panic("打开桶b1失败")
		}
		v1 := bucket.Get([]byte("1"))
		v2 := bucket.Get([]byte("2"))
		fmt.Printf("v1: %s \nv2: %s\n",v1,v2)
		return nil
	})
}
