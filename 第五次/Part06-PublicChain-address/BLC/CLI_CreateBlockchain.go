package BLC

func (cli *CLI)CreateBlockchain(address string)  {
	//fmt.Println("创建创世区块...")

	CreateBlockchainWithGenesisBlock(address)
}
