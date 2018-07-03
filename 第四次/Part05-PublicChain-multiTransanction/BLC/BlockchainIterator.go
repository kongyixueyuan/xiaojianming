package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

type BlockchainIterator struct{
	CurrentHash []byte
	DB *bolt.DB
}

//迭代器，获取迭代器对象
func (blockchain *BlockChain)Iterator()*BlockchainIterator{
	return &BlockchainIterator{blockchain.Tip,blockchain.DB}
}

//迭代器Next函数,返回Block对象，修改迭代器当前CurrentHash
func (it *BlockchainIterator)Next()*Block{
	var block *Block
	err := it.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TableName))
		if b != nil{
			block = DeSerialize(b.Get(it.CurrentHash))
			it.CurrentHash = block.PrevHash
		}
		return nil
	})
	if err != nil{
		log.Panic("Next()：迭代器view数据表失败", err)
	}
	return block
}
