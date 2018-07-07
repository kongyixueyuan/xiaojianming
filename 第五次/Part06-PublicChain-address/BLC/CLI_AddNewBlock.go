package BLC

import (
	"fmt"
	"os"
)

//添加区块到区块链
func (cli *CLI)AddNewBlock(txs []*Transaction){
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
