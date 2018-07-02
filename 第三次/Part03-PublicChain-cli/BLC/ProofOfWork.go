package BLC

import (
	"math/big"
	"bytes"
	"math"
	"crypto/sha256"
	"fmt"
)

const DiffcultyBit int = 16
const MaxNonce int64 = math.MaxInt64

type ProofOfWork struct{
	block *Block
	TargetBit *big.Int
}

func NewProofOfWork(block *Block) *ProofOfWork{
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - DiffcultyBit))
	return &ProofOfWork{block, target}
}

func (pow *ProofOfWork)PrepareData(nonce int64)[]byte {

	data := bytes.Join([][]byte{
		IntToHex(pow.block.Timestamp),
		IntToHex(pow.block.Height),
		pow.block.PrevHash,
		IntToHex(nonce)},
		[]byte{})

	return data
}

func (pow *ProofOfWork)Run(){
	var hashInt big.Int
	var nonce int64 = 0

	for nonce < MaxNonce{
		hash := sha256.Sum256(pow.PrepareData(nonce))
		hashInt.SetBytes(hash[:])
		fmt.Printf("\r0x%x", hash)

		if hashInt.Cmp(pow.TargetBit) == -1{
			pow.block.Hash = hash[:]
			break
		}
		nonce++
	}
	fmt.Printf("\n")
}