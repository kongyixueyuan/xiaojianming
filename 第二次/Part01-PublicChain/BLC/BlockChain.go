package BLC




type BlockChain struct {
	Blocks []*Block
}

//创建区块链
func CreateBlockChain()*BlockChain{
	gensisBlock := CreateGenesisBlock()
	return &BlockChain{[]*Block{gensisBlock}}
}

//添加新的区块
func (bc *BlockChain)AddNewBlock(data string){
	currentHeigth := int64(len(bc.Blocks))
	block := CreateNewBlock(data, currentHeigth + 1, bc.Blocks[currentHeigth -1].Hash)
	bc.Blocks = append(bc.Blocks, block)
}