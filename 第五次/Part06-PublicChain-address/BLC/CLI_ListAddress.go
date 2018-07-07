package BLC

import (
	"fmt"
	"log"
)

func (cli *CLI)CLI_ListAddress(){
	fmt.Println("钱包地址清单：")
	ws,err := NewWallets()
	if err != nil {
		log.Panic(err)
	}
	addresses := ws.GetAddress()

	for _, address := range addresses {
		fmt.Println(address)
	}
}