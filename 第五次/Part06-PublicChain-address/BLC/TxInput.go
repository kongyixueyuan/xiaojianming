package BLC

import "bytes"

type TxInput struct{
	TxHash []byte
	//存储TxOutput在Vout里面的索引
	Vout int64
	//用户明
	//ScriptSig string

	Signature []byte
	PubKey    []byte
}


func (in *TxInput) CanUnlockOutputWith(ripemd160hash []byte) bool {
	pubKey := HashPupKey(in.PubKey)
	return bytes.Compare(pubKey, ripemd160hash) == 0
}

func (in *TxInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPupKey(in.PubKey)
	return bytes.Compare(lockingHash, pubKeyHash) == 0
}