package main

import (
	"fmt"
	"os"
)

//接收命令行参数,控制区块链操作
type CLI struct {
	bc *BlockChain
}
const Usage=`
	addBlock --data DATA "add data to BlockChain"
	printChain "print all blockChain data"
`
func (cli *CLI)Run()  {
	// 1. 得到命令
	args := os.Args
	if len(args)<2{
		fmt.Printf(Usage)
		return
	}
	// 2. 分析命令
	cmd:=args[1]
	switch cmd {
	case "addBlock":
		if len(args)==4&&args[2]=="--data"{
			// 添加区块
			// a. 获取数据
			cli.AddBlock(args[3])
		}else{
			fmt.Println("添加区块失败,参数错误,请检查.")
			fmt.Printf(Usage)
		}
		// b. 使用bc添加区块到AddBlock
	case "printChain":
		// 打印区块
		cli.PrintBlockChain()
	default:
		fmt.Println("无效命令,请检查.")
		fmt.Printf(Usage)
	}
	// 3. 执行相应操作
}
