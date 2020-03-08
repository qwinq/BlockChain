package main

import "fmt"

func (cli *CLI)PrintBlockChain()  {
	it := cli.bc.NewBcIterator()
	for {
		block := it.Next()
		fmt.Println("=================================")
		fmt.Printf("版本号:%d\n", block.Version)
		fmt.Printf("前区块哈希值:%x\n", block.PrevHash)
		fmt.Printf("梅克尔根:%x\n", block.MerkleRoot)
		fmt.Printf("时间戳:%d\n", block.TimeStamp)
		fmt.Printf("难度值(待改):%d\n", block.Difficulty)
		fmt.Printf("随机数:%d\n", block.Nonce)
		fmt.Printf("本区块哈希值:%x\n", block.Hash)
		fmt.Printf("区块数据:%s\n", string(block.Data))
		if len(block.PrevHash)==0{
			fmt.Printf("遍历结束!\n")
			break
		}
	}
}
func (cli *CLI)AddBlock(data string)  {
	cli.bc.AddBlock(data)
	fmt.Printf("添加区块成功!\n")
}
