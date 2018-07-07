package BLC

import (
	"os"
	"io/ioutil"
	"log"
	"encoding/gob"
	"crypto/elliptic"
	"bytes"
)

const walletFile = "wallet.dat"

//具有多个地址的钱包
type Wallets struct{
	Wallets map[string]*Wallet
}

func NewWallets()(*Wallets,error){

	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)

	err := wallets.LoadFromFile()
	return &wallets, err
}

func (ws *Wallets)LoadFromFile()error{
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	ws.Wallets = wallets.Wallets
	return nil
}

func (ws *Wallets)CreateWallet()string{
	wallet := NewWallet()
	address := wallet.GetAddress()
	ws.Wallets[string(address)] = wallet
	return string(address)
}

func (ws Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}

func (ws *Wallets)GetAddress()[]string{
	var addresses []string
	for address := range ws.Wallets{
		addresses = append(addresses,address)
	}
	return addresses
}