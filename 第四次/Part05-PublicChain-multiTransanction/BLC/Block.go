package BLC

import (
	"time"
	"bytes"
	"fmt"
	"encoding/gob"
	"log"
	"crypto/sha256"
)

type Block struct{
	// 1时间戳
	Timestamp int64
	//2 高度
	Height int64
	// data
	//Data []byte
	Txs []*Transaction
	//3、前一个区块的hash
	PrevHash []byte
	//4、当前区块的hash
	Hash []byte
	//5、随机值nonce
	Nonce int64
}

//创建一个区块
func CreateNewBlock(txs []*Transaction, height int64, prevHash []byte) *Block{

	block := &Block{
		time.Now().Unix(),
		height,
		txs,
		prevHash,
		nil,
		0}
	//block.SetHash()
	pow := NewProofOfWork(block)
	pow.Run()
	return block
}
//创建创始区块
func CreateGenesisBlock(tx []*Transaction) *Block{

	return CreateNewBlock(tx, 1, []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
//打印区块信息
func (block *Block)GetBlockInfo(){
	fmt.Println("[")
	fmt.Println("Block Height:", block.Height)
	fmt.Printf("Txs:\n")
	for _,tx:=range block.Txs{
		fmt.Printf("Transaction Hash:0x%x\n", tx.TxHash)
		fmt.Println("Vins:")
		for _,in := range tx.Vins{
			fmt.Printf("[0x%x %d, %s]\n", in.TxHash, in.Vout, in.ScriptSig)
		}

		fmt.Println("Vouts:")
		for _,out := range tx.Vouts{
			fmt.Printf("[%d ", out.Value)
			fmt.Printf("%s]\n", out.ScriptPubKey)
		}
	}
	fmt.Println("Timestamp:", time.Unix(block.Timestamp,0).Format("2006-01-02 03:04:05 PM"))
	fmt.Printf("PrevHash:0x%x\n", block.PrevHash)
	fmt.Printf("Hash    :0x%x\n", block.Hash)
	fmt.Printf("Nonce:%d\n", block.Nonce)
	fmt.Println("]")
}

//序列化区块,将区块序列化字节数组
func (block *Block)Serialize()[]byte{
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	if (err != nil){
		fmt.Println(err)
		log.Panic(err)
	}
	return result.Bytes()
}

//反序列化
func DeSerialize(data []byte)*Block{
	//var result bytes.Buffer
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))

	err := decoder.Decode(&block)
	if err != nil{
		fmt.Println(err)
		log.Panic(err)
	}
	return &block
}

func (block *Block)HashTransactions()[]byte{

	var txHashes [][]byte
	var txHash [32]byte
	for _,tx:= range block.Txs{
		txHashes = append(txHashes, tx.TxHash)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}


