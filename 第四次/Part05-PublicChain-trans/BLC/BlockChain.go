package BLC

import (
	"github.com/boltdb/bolt"
	"os"
	"fmt"
	"log"
	"math/big"
	"encoding/hex"
)

const DBName = "blockchain.db"
const TableName = "blocks"

type BlockChain struct {
	Tip []byte //最新区块的hash
	DB *bolt.DB
}

//获取区块链对象
func GetBlockchainObject()*BlockChain{
	//打开数据库
	db, err := bolt.Open(DBName, 0600,nil)
	if err != nil{
		log.Panic("GetBlockchainObject():打开数据库失败", err)
	}
	var tip []byte //存储区块链上最新区块的hash
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TableName))
		if b!=nil{
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	return &BlockChain{tip, db}
}

// 创建带有创世区块的区块链
func CreateBlockchainWithGenesisBlock(address string) *BlockChain{
	if DBExist(){
		fmt.Println("创世区块已经存在...")
		os.Exit(1)
	}
	fmt.Println("正在创建创世区块......")
	var bc BlockChain
	//创建数据库
	db, err := bolt.Open(DBName, 0600, nil)
	defer db.Close()
	if err!=nil{
		log.Panic("创建创世区块失败", err)
	}
	//更新数据，添加创世区块到数据库
	err = db.Update(func(tx *bolt.Tx) error {
		//创建表
		b,err := tx.CreateBucket([]byte(TableName))
		//创建创世区块
		txConbase := NewCoinBaseTransaction(address)
		genesisblock := CreateGenesisBlock([]*Transaction{txConbase})
		err = b.Put(genesisblock.Hash, genesisblock.Serialize())
		if err != nil{
			log.Panic("添加创世区块失败", err)
		}
		bc.Tip = genesisblock.Hash
		//存储最新区块的Hash
		err = b.Put([]byte("l"),genesisblock.Hash)
		if err != nil{
			log.Panic("添加最新区块Hash失败", err)
		}
		return nil
	})
	if err != nil{
		log.Panic("添加创世区块到数据库失败", err)
	}
	bc.DB = db
	return &bc
}

// 判断数据库是否存在
func DBExist()bool{
	if _,err := os.Stat(DBName);os.IsNotExist(err){
		return false
	}
	return true
}

//遍历Blockchain
func (bc BlockChain)PrintBlockchain(){
	bit := bc.Iterator()
	for{
		block := bit.Next()
		block.GetBlockInfo()
		var hashInt big.Int
		hashInt.SetBytes(block.PrevHash)
		if hashInt.Cmp(big.NewInt(0)) == 0{
			break
		}
	}
}

//添加新的区块到区块链
func (bc *BlockChain)AddBlockToBlockchain(txs []*Transaction){
	//更细数据库
	err := bc.DB.Update(func(tx *bolt.Tx) error {

		b := tx.Bucket([]byte(TableName))
		if b != nil{

			//获取最新区块
			block := DeSerialize(b.Get(bc.Tip))

			//创建新的区块
			newBlock := CreateNewBlock(txs, block.Height + 1,block.Hash)
			err := b.Put(newBlock.Hash,newBlock.Serialize())
			if err != nil{
				log.Panic("AddBlockToBlockchain():Put最新区块到数据库失败")
			}
			bc.Tip = newBlock.Hash
			b.Put([]byte("l"), bc.Tip)
			if err != nil{
				log.Panic("AddBlockToBlockchain():Put最新区块Hash到数据库失败")
			}
		}
		return nil
	})
	if err!= nil{
		log.Panic("AddBlockToBlockchain():Update数据库失败",err)
	}
}
//挖矿产生新区块
func (bc *BlockChain)MineNewBlock(from, to, amount []string){
	//fmt.Println(from)
	//fmt.Println(to)
	//fmt.Println(amount)
	var txs []*Transaction
	var block *Block
	err := bc.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TableName))
		block  = DeSerialize(b.Get(bc.Tip))
		return nil
	})
	if err != nil{
		log.Panic("MineNewBlock():查询数据库失败")
	}
	err = bc.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TableName))
		if b == nil{
			log.Panic("数据库表为空")
		}
		newBlock := CreateNewBlock(txs,block.Height + 1, block.Hash)
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err !=nil{
			log.Panic("插入新的区块失败", err)
		}
		err = b.Put([]byte("l"),newBlock.Hash)
		if err !=nil{
			log.Panic("插入最新区块的hash失败", err)
		}
		bc.Tip = newBlock.Hash
		return nil
	})
	if err != nil{
		log.Panic("MineNewBlock():更新数据库失败")
	}
}
////创建区块链
//func CreateBlockChain()*BlockChain{
//	gensisBlock := CreateGenesisBlock()
//	return &BlockChain{[]*Block{gensisBlock}}
//}
//
////添加新的区块
//func (bc *BlockChain)AddNewBlock(data string){
//	currentHeigth := int64(len(bc.Blocks))
//	block := CreateNewBlock(data, currentHeigth + 1, bc.Blocks[currentHeigth -1].Hash)
//	bc.Blocks = append(bc.Blocks, block)
//}

//找到未花费输出的交易
func (bc *BlockChain)FindUnspendTransaction(address string)[]Transaction{

	var unspentTXs []Transaction
	spentTXOs := make(map[string][]int)
	bci := bc.Iterator()
	for {
		block := bci.Next()
		for _, tx := range block.Txs{
			txHash := hex.EncodeToString(tx.TxHash)

			//1、遍历输出
			Outputs:
			for outIdx,out:= range tx.Vouts {
				if spentTXOs[txHash] != nil {
					for _,spendOut := range spentTXOs[txHash]{
						if spendOut == outIdx{
							continue Outputs
						}
					}
				}
				if out.CanBeUnlockedWith(address){
					unspentTXs = append(unspentTXs, *tx)
				}
			}
			if tx.IsCoinBaseTx() == false {
				for _, in := range tx.Vins{
					if in.CanUnlockOutputWith(address) {
						inTxID := hex.EncodeToString(in.TxHash)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Vout)
					}
				}
			}
		}
		//遍历到头退出循环
		var hashInt big.Int
		hashInt.SetBytes(block.PrevHash)
		if hashInt.Cmp(big.NewInt(0)) == 0{
			break
		}
	}
	return unspentTXs
}
