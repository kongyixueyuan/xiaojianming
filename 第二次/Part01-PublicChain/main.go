package main

import (
	"./BLC"
)
func main(){
	bc := BLC.CreateBlockChain()
	bc.AddNewBlock("A send 10 BTM to B")
	bc.AddNewBlock("A send 1 BTC to C")
	//bc.Blocks[0].GetBlockInfo()
	//bc.Blocks[1].GetBlockInfo()
	//bc.Blocks[2].GetBlockInfo()
	//fmt.Println(bc.Blocks[0].Serialize())
	BLC.DeSerialize(bc.Blocks[0].Serialize()).GetBlockInfo()
	BLC.DeSerialize(bc.Blocks[1].Serialize()).GetBlockInfo()
	BLC.DeSerialize(bc.Blocks[2].Serialize()).GetBlockInfo()

}

func CreateBoltdbTest(){
	
}