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

//send
func (cli *CLI)Send(from, to ,amount []string){
	cli.Blockchain  = GetBlockchainObject()
	defer cli.Blockchain.DB.Close()

	cli.Blockchain.MineNewBlock(from, to ,amount)
}
func (cli *CLI)Run() {
	IsValidArgs()

	createBlockchainCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printblockchain", flag.ExitOnError)

	createBlockchainData := createBlockchainCmd.String("address", "", "创建创世区块的地址")
	fromData := sendCmd.String("from", "", "交易源地址")
	toData := sendCmd.String("to", "", "交易目的地址")
	amountData := sendCmd.String("amount", "", "交易数量")

	switch os.Args[1] {
		case "createblockchain":
			err := createBlockchainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "send":
			err := sendCmd.Parse(os.Args[2:])
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
		if (*createBlockchainData == ""){
			fmt.Println("地址不能为空")
			PrintUsage()
			os.Exit(1)
		}
		CreateBlockchainWithGenesisBlock(*createBlockchainData)
	}
	if sendCmd.Parsed(){
		//fmt.Println("add 区块...")
		if *fromData == "" || *toData =="" || *amountData==""{
			PrintUsage()
			os.Exit(1)
		}

		//fmt.Println(s)
		//fmt.Println(*toData)
		//fmt.Println(*amountData)
		cli.Send(JSONToArray(*fromData),JSONToArray(*toData),JSONToArray(*amountData))
		//fmt.Println(JSONToArray(*fromData))
		//fmt.Println(JSONToArray(*toData))
		//fmt.Println(JSONToArray(*amountData))
		//cli.addNewBlock([]*Transaction{})
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
func (cli *CLI)addNewBlock(txs []*Transaction){
	if 	!DBExist(){
		fmt.Printf("数据库不存在")
		os.Exit(1)
	}
	//获取区块链对象
	cli.Blockchain = GetBlockchainObject()
	defer cli.Blockchain.DB.Close()
	cli.Blockchain.AddBlockToBlockchain(txs)
	//defer bc.DB.Close()
	//bc.AddBlockToBlockchain(data)
}

//打印命令行提示信息
func PrintUsage(){
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -address DATA --创建区块链")
	fmt.Println("\tsend -form FROM -to TO -amount AMOUNT --交易明细")
	fmt.Println("\tprintblockchain --打印区块链信息")
}

//输入参数是否有效
func IsValidArgs(){
	if len(os.Args) < 2{
		PrintUsage()
		os.Exit(1)
	}
}
