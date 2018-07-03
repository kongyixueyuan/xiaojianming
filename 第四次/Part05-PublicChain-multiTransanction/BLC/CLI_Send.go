package BLC

//send
func (cli *CLI)Send(from, to ,amount []string){
	cli.Blockchain  = GetBlockchainObject()
	defer cli.Blockchain.DB.Close()

	cli.Blockchain.MineNewBlock(from, to ,amount)
}
