package main

import "log"

func main() {
	bc := NewBlockChain()

	bc.AddBlock("Send 1 BTC to IVan")
	bc.AddBlock("Send 2 more BTC to IVan")

	for _, block := range bc.blocks {
		log.Printf("Prev. hash :%x\n", block.PreBlockHash)
		log.Printf("Data: %s\n", block.Data)
		log.Printf("Hash: %x\n", block.Hash)
		log.Println()
	}
}
