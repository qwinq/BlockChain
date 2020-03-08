package main

//TODO 简单版
// 1. 定义结构  1. 前区块哈希 2. 当前区块哈希 3. 数据
// 2. 创建区块 NewBlock(data string, prevBlockHash []byte) *Block
// 3. 生成哈希  1. 拼接数据 2. 生成哈希 SetHash()
// 4. 引入区块链 BlockChain struct
// 5. 定义区块链 NewBlockChain() *BlockChain 1. GenesisBlock() *Block
// 6. 添加区块 (bc *BlockChain)AddBlock(data string) a.创建新的区块 b. 添加到区块链数组中
// 7. 重构代码 分文件(分包)

//TODO 升级版
// 1. 补充区块字段
// 2. 更新计算哈希函数 sha256.Sum256(blockInfo)
// 3. 优化代码 bytes.Join(tmp, []byte(""))
func main() {
	bc := NewBlockChain()
	cli:=CLI{bc:bc}
	cli.Run()
	//bc.AddBlock("Second Block")
	//bc.AddBlock("Third Block")

	//for i,block:=range bc.blocks{
	//	fmt.Printf("======当前高度:%d======\n", i)
	//	fmt.Printf("前区块哈希值:%x\n", block.PrevHash)
	//	fmt.Printf("当前区块哈希值:%x\n", block.Hash)
	//	fmt.Printf("数据:%s\n", string(block.Data))
	//}
	/*it := bc.NewBcIterator()
	for  {
		block:=it.Next()

		fmt.Println("===================")
		fmt.Printf("前区块哈希值:%x\n", block.PrevHash)
		fmt.Printf("当前区块哈希值:%x\n", block.Hash)
		fmt.Printf("区块数据:%s\n", string(block.Data))
		if block.PrevHash==nil{
		//if len(block.PrevHash)==0{
			break
		}
	}
	*/
}
//TODO V3
// v3版本思路
// 1. BlockChain结构重写 使用数据库代替数组
// 2. NewBlockChain函数重写 由对数组操作改写成==对数据库操作==，创建数据库
// 3. AddBlock函数重写 对数据库的读取和写入
// 4.打印数据 对数据库的遍历（迭代器Iterator）
// 5.命令行 1.添加区块命令 2.打印区块链命令 3.创建区块链（可选）