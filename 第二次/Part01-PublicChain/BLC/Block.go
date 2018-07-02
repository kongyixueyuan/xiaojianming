package BLC

import (
	"time"
	"strconv"
	"bytes"
	"crypto/sha256"
	"fmt"
	"encoding/gob"
	"log"
)

type Block struct{
	// 1时间戳
	Timestamp int64
	//2 高度
	Height int64
	// data
	Data []byte
	//3、前一个区块的hash
	PrevHash []byte
	//4、当前区块的hash
	Hash []byte
	//5、随机值nonce
	Nonce int64
}

//创建一个区块
func CreateNewBlock(data string, height int64, prevHash []byte) *Block{

	block := &Block{
		time.Now().Unix(),
		height,[]byte(data),
		prevHash,
		nil,
		0}
	//block.SetHash()
	pow := NewProofOfWork(block)
	pow.Run()
	return block
}
//创建创始区块
func CreateGenesisBlock() *Block{
	return CreateNewBlock("Genesis Block...", 0, []byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0})
}
// setHash，设置区块的hash值
func (b *Block)SetHash() {
	heightBytes := IntToHex(b.Height)
	timeString := strconv.FormatInt(b.Timestamp, 2)
	timeBytes := []byte(timeString)
	nonceBytes := IntToHex(b.Nonce)
	blockBytes := bytes.Join([][]byte{timeBytes, heightBytes, b.Data, b.PrevHash, b.Hash, nonceBytes}, []byte{})

	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}

//打印区块信息
func (block *Block)GetBlockInfo(){
	fmt.Println("[")
	fmt.Println("Block Height:", block.Height)
	fmt.Println("Transaction data:", string(block.Data))
	fmt.Println("Timestamp:", block.Timestamp)
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