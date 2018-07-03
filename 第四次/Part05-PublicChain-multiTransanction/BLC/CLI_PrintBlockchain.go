package BLC

import (
	"fmt"
	"os"
)

//打印区块链
func (cli *CLI)PrintBlockchain(){
	if 	!DBExist(){
		fmt.Printf("数据库不存在")
		os.Exit(1)
	}
	//获取区块链对象
	cli.Blockchain = GetBlockchainObject()
	defer cli.Blockchain.DB.Close()
	if cli.Blockchain == nil{
		fmt.Println("区块链不存在...")
		PrintUsage()
		os.Exit(1)
	}
	cli.Blockchain.PrintBlockchain()
}
