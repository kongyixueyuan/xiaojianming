package BLC

import (
	"os"
	"fmt"
	"flag"
	"log"
)

type CLI struct{
	Blockchain *BlockChain
}
func (cli *CLI)Run() {
	IsValidArgs()

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printblockchain", flag.ExitOnError)

	addBlockData := addBlockCmd.String("data", "交易数据", "增加交易数据")

	switch os.Args[1] {
		case "createblockchain":
			err := createBlockchainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "addblock":
			err := addBlockCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "printblockchain":
			err := printChainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		default:
			PrintUsage()
			os.Exit(1)
	}
	if createBlockchainCmd.Parsed(){
		//fmt.Println("创建创世区块...")
		CreateBlockchainWithGenesisBlock()
	}
	if addBlockCmd.Parsed(){
		//fmt.Println("add 区块...")
		cli.addNewBlock(*addBlockData)
	}
	if printChainCmd.Parsed(){
		cli.PrintBlockchain()
	}
}

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

//添加区块到区块链
func (cli *CLI)addNewBlock(data string){
	if 	!DBExist(){
		fmt.Printf("数据库不存在")
		os.Exit(1)
	}
	//获取区块链对象
	cli.Blockchain = GetBlockchainObject()
	defer cli.Blockchain.DB.Close()
	cli.Blockchain.AddBlockToBlockchain(data)
	//defer bc.DB.Close()
	//bc.AddBlockToBlockchain(data)
}

//打印命令行提示信息
func PrintUsage(){
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain --创建区块链")
	fmt.Println("\taddblock -data DATA --	增加区块")
	fmt.Println("\tprintblockchain --打印区块链信息")
}

//输入参数是否有效
func IsValidArgs(){
	if len(os.Args) < 2{
		PrintUsage()
		os.Exit(1)
	}
}
