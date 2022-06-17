package main

import (
	"fmt"
	"log"
)

func (cli *CLI) createBlockChain(address string) {
	if !ValidateAddress(address) {
		log.Panic("ERROR: Address if not valid")
	}

	bc := CreateBlockChain(address)
	bc.db.Close()
	fmt.Println("Done!")
}
