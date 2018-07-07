package main

import (
	"crypto/sha256"
	"fmt"
	"./BLC"
	"github.com/ethereum/go-ethereum/vendor/golang.org/x/crypto/ripemd160"
)

func main(){

	cli:= BLC.CLI{}
	cli.Run()

}

//sha256
func sha256Test(){
	haser := sha256.New()
	haser.Write([]byte("xiaojianmig"))
	bytes := haser.Sum(nil)
	fmt.Printf("0x%x\n", bytes)

}

func base58Test(){
	bytes := []byte("xiaojianming")
	bytes58 := BLC.Base58Encode(bytes)
	fmt.Printf("%x\n", bytes58)
	fmt.Printf("%s\n", bytes58)
	bytes = BLC.Base58Decode(bytes58)
	fmt.Printf("%s", bytes)
}

func ripemd160Test(){
	haser := ripemd160.New()
	haser.Write([]byte("xiaojianmig"))
	bytes := haser.Sum(nil)
	fmt.Printf("0x%x\n", bytes)

}