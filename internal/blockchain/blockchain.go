package blockchain

import "github.com/IvanKalug-QA/Blockchain-on-Go/internal/blocks"

type Blockchain struct {
	Blocks []*blocks.Block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlockHash := bc.Blocks[len(bc.Blocks)-1]
	newBlock := blocks.NewBlock(data, prevBlockHash.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*blocks.Block{blocks.NewGenesisBlock()}}
}
