package BLC

import "bytes"

type TxOutput struct {
	Value int64			//一定量的币
	//ScriptPubKey string	//用户

	PubKeyHash []byte
}


func (out *TxOutput) CanBeUnlockedWith(unlockingData string) bool {

	publicKeyHash := Base58Decode([]byte(unlockingData))
	hash160 := publicKeyHash[1:len(publicKeyHash)-addressCheckSumLen]

	return bytes.Compare(out.PubKeyHash, hash160) == 0
}

func NewTxOutput(value int64, address string)*TxOutput{
	txOutput := &TxOutput{value, nil}

	txOutput.Lock(address)
	return txOutput
}

func (txo *TxOutput)Lock(address string){

	pubkeyHash := Base58Decode([]byte(address))
	pubkeyHash = pubkeyHash[1:len(pubkeyHash)-addressCheckSumLen]
	txo.PubKeyHash = pubkeyHash
}