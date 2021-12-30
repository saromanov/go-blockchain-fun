package blockchain

type Block struct {
	Hash []byte
	Data []byte
	Prev []byte
	Nonce int
}

func CreateBlock(data, prevHash []byte) *Block {
    block := &Block{[]byte{}, data, prevHash}
    block.Hashing()
    return block
}

func (b *Block) Hashing() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:] 
  }

type BlockChain struct {
	Blocks []*Block
}

func (chain *BlockChain) AddBlock(data []byte) {
    prevBlock := chain.blocks[len(chain.blocks)-1]
    new := CreateBlock(data, prevBlock.Hash)
    chain.blocks = append(chain.blocks, new)
}