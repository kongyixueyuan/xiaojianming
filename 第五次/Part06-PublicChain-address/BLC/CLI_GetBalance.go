package BLC

import "fmt"

//命令行查询余额
func (cli *CLI)GetBalance(address string){
	cli.Blockchain = GetBlockchainObject()
	defer cli.Blockchain.DB.Close()

	//获取未花费输出
	//txOuts := cli.Blockchain.UTXOOut(address)
	//txOuts := cli.Blockchain.FindUTXO(address)
	//var balances int64 = 0
	//for _, out := range txOuts{
	//	balances += out.Value
	//}
	balances := cli.Blockchain.GetBalance(address)
	fmt.Printf("Balances of %s:Tokens:%d\n", address,balances)
}

