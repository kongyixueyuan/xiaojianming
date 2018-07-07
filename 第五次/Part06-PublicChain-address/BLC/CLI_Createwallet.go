package BLC

import "fmt"

//创建钱包
func (cli *CLI)CreateWallet(){

	wallets , _ := NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	fmt.Printf("你的新地址是:%s", address)

}

//判断地址是否有效
func (cli *CLI)CLI_IsValidAddress(address string){
	fmt.Println(address,"is", IsValidAddress([]byte(address)))
}