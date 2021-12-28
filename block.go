package blockchain

type Block struct {
	Hash []byte
	Data []byte
	Prev []byte
	Nonce int
}

func (b *Block) Hashing() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:] 
  }

type BlockChain struct {
	Blocks []*Block
}