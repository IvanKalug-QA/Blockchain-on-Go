package main

import (
	"fmt"

	"github.com/IvanKalug-QA/Blockchain-on-Go/internal/blockchain"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddBlock("Send 1 BTC from Ivan")
	bc.AddBlock("Send 0.2 BTC from Jhon")

	for _, block := range bc.Blocks {
		fmt.Printf("Prev. Hash %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
