package blockchain

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data []byte) {
    prevBlock := chain.blocks[len(chain.blocks)-1]
    new := CreateBlock(data, prevBlock.Hash)
    chain.blocks = append(chain.blocks, new)
}

func Create()*BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}