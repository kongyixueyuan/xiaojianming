package BLC

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"crypto/sha256"
	"math/big"
)

type Transaction struct{

	//1、交易Hash
	TxHash []byte

	//2、输入
	Vins []*TxInput
	//3、输出
	Vouts []*TxOutput
}

//创建Transaction
//1、创建创世区块的Transaction
func NewCoinBaseTransaction(address string)*Transaction{

	txInput := &TxInput{[]byte{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0}, -1, "Genesis block..."}
	txOutput := &TxOutput{10, address}

	txCoinbase := &Transaction{[]byte{},[]*TxInput{txInput}, []*TxOutput{txOutput} }
	//设置Hash值
	txCoinbase.HashTransaction()
	return txCoinbase
}

//2、转账时产生的Transaction
func NewSimpleTransaction(from, to string, amount int64)*Transaction{

	return nil
}


func (tx *Transaction)HashTransaction(){
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(tx)
	if (err != nil){
		fmt.Println(err)
		log.Panic(err)
	}
	hash := sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}

//
func (bc *BlockChain)FindUnSpendTransaction(address string)*Transaction{

	return nil
}

//1、找到未花费的输出
func (bc *BlockChain)FindSpendableTxOupts(address string, amount int64)(int, map[string][]int64	){
	//unspentOutputs := make(map[string][]int64)

	return 0, nil
}

//判断是否是coinbase transaction
func (tx *Transaction)IsCoinBaseTx()bool{
	var hashInt big.Int
	hashInt.SetBytes(tx.TxHash)
	///判断交易hash是否等0
	if hashInt.Cmp(big.NewInt(0)) == 0{
		return true
	}
	return false
}