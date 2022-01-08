package blockchain

type BlockChain struct {
	Blocks []*Block
}

// AddBlock provides adding of the new block to blockchain
func (chain *BlockChain) AddBlock(data []byte) {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    new := CreateBlock(data, prevBlock.Hash)
    chain.Blocks = append(chain.Blocks, new)
}

// Create provides creating of the new block with genesis
func Create()*BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}