package blockchain

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data []byte) {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    new := CreateBlock(data, prevBlock.Hash)
    chain.Blocks = append(chain.Blocks, new)
}

func Create()*BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}