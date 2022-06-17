package main

import "log"

func (cli *CLI) createWallet() {
	wallets, _ := NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()
	log.Printf("Your new address: %s\n", address)
}
