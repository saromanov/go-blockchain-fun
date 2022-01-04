package blockchain

type Block struct {
	Hash []byte
	Data []byte
	Prev []byte
	Nonce int
}

func CreateBlock(data, prevHash []byte) *Block {
    block := &Block{
		Hash: []byte{}, 
		Data:data, 
		Prev: prevHash
	}
    block.Hashing()
    return block
}

func Genesis() *Block {
	return CreateBlock([]byte("Genesis"), []byte{})
}

func (b *Block) Hashing() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:] 
  }