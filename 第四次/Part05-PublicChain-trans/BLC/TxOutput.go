package BLC

type TxOutput struct {
	Value int64			//一定量的币
	ScriptPubKey string	//用户
}

func (out *TxOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}