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
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printblockchain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)

	createBlockchainData := createBlockchainCmd.String("address", "", "创建创世区块的地址")
	fromData := sendCmd.String("from", "", "交易源地址")
	toData := sendCmd.String("to", "", "交易目的地址")
	amountData := sendCmd.String("amount", "", "交易数量")
	getBalanceAddr := getBalanceCmd.String("address", "", "获取账户余额")

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
		case "getBalance":
			err := getBalanceCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		default:
			PrintUsage()
			os.Exit(1)
	}
	if createBlockchainCmd.Parsed(){
		if (*createBlockchainData == ""){
			fmt.Println("地址不能为空")
			PrintUsage()
			os.Exit(1)
		}
		cli.CreateBlockchain(*createBlockchainData)
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

	if getBalanceCmd.Parsed(){
		if *getBalanceAddr == ""{
			fmt.Println("地址不能为空......")
			PrintUsage()
			os.Exit(1)
		}
		cli.GetBalance(*getBalanceAddr)
	}
}

//打印命令行提示信息
func PrintUsage(){
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -address DATA --创建区块链")
	fmt.Println("\tsend -form FROM -to TO -amount AMOUNT --交易明细")
	fmt.Println("\tprintblockchain --打印区块链信息")
	fmt.Println("\tgetBalance -address DATA --查询账户余额")
}

//输入参数是否有效
func IsValidArgs(){
	if len(os.Args) < 2{
		PrintUsage()
		os.Exit(1)
	}
}
