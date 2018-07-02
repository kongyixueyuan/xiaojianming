package BLC

type TxInput struct{
	TxHash []byte
	//存储TxOutput在Vout里面的索引
	Vout int
	//用户明
	ScriptSig string
}
func (in *TxInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}
