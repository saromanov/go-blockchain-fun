package blockchain

import "errors"

type BlockChain struct {
	Blocks []*Block
}

// AddBlock provides adding of the new block to blockchain
func (chain *BlockChain) AddBlock(data []byte) error  {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    newBlock := CreateBlock(data, prevBlock.Hash)
    if ok := IsBlockValid(prevBlock, newBlock); ok {
        return errors.New("unable to validate blocks")
    }
    chain.Blocks = append(chain.Blocks, newBlock)
}

// Create provides creating of the new block with genesis
func Create()*BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}