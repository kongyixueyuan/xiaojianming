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
	createWalletCmd := flag.NewFlagSet("createWallet", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printChain", flag.ExitOnError)
	getBalanceCmd := flag.NewFlagSet("getBalance", flag.ExitOnError)
	isValidAddressCmd := flag.NewFlagSet("isValidAddress", flag.ExitOnError)
	listAddressCmd := flag.NewFlagSet("listAddress", flag.ExitOnError)

	createBlockchainData := createBlockchainCmd.String("address", "", "创建创世区块的地址")
	fromData := sendCmd.String("from", "", "交易源地址")
	toData := sendCmd.String("to", "", "交易目的地址")
	amountData := sendCmd.String("amount", "", "交易数量")
	getBalanceAddr := getBalanceCmd.String("address", "", "获取账户余额")
	isValidAddressAddr := isValidAddressCmd.String("address", "", "区块链地址")

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
		case "printChain":
			err := printChainCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "getBalance":
			err := getBalanceCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "createWallet":
			err := createWalletCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "isValidAddress":
			err := isValidAddressCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		case "listAddress":
			err := listAddressCmd.Parse(os.Args[2:])
			if err != nil {
				log.Panic(err)
			}
		default:
			PrintUsage()
			os.Exit(1)
	}
	if createBlockchainCmd.Parsed(){
		fmt.Println(*createBlockchainData)
		if (IsValidAddress([]byte(*createBlockchainData))) == false{
			fmt.Println("地址无效......")
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
		from := JSONToArray(*fromData)
		to := JSONToArray(*toData)
		for index, addr := range from{
			if IsValidAddress([]byte(addr)) == false || IsValidAddress([]byte(to[index])) == false{
				fmt.Println("地址无效......")
				PrintUsage()
				os.Exit(1)
			}
		}
		cli.Send(from,to,JSONToArray(*amountData))

	}
	if printChainCmd.Parsed(){
		cli.PrintBlockchain()
	}

	if getBalanceCmd.Parsed(){
		if IsValidAddress([]byte(*getBalanceAddr)) == false{
			fmt.Println("无效地址......")
			PrintUsage()
			os.Exit(1)
		}
		cli.GetBalance(*getBalanceAddr)
	}

	if createWalletCmd.Parsed(){
		cli.CreateWallet()
	}

	if isValidAddressCmd.Parsed(){
		cli.CLI_IsValidAddress(*isValidAddressAddr)
	}

	if listAddressCmd.Parsed(){
		cli.CLI_ListAddress()
	}
}

//打印命令行提示信息
func PrintUsage(){
	fmt.Println("Usage:")
	fmt.Println("\tcreateblockchain -address DATA --创建区块链")
	fmt.Println("\tcreateWallet --创建区块链")
	fmt.Println("\tsend -form FROM -to TO -amount AMOUNT --交易明细")
	fmt.Println("\tprintChain --打印区块链信息")
	fmt.Println("\tgetBalance -address DATA --查询账户余额")
	fmt.Println("\tisValidAddress -address DATA --判断地址是否有效")
	fmt.Println("\tlistAddress --打印钱包地址")
}

//输入参数是否有效
func IsValidArgs(){
	if len(os.Args) < 2{
		PrintUsage()
		os.Exit(1)
	}
}
