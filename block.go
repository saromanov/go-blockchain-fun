package blockchain

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	ID        int64
	Hash      []byte
	Data      []byte
	Prev      []byte
	Nonce     int
	Timestamp int64
}

// CreateBlock provides initialization of block
func CreateBlock(data, prevHash []byte) *Block {
	block := &Block{
		ID:        0,
		Hash:      []byte{},
		Data:      data,
		Prev:      prevHash,
		Timestamp: time.Now().UTC().Unix(),
	}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

// Genesis provides initalization of blockchain
func Genesis() *Block {
	return CreateBlock([]byte("Genesis"), []byte{})
}

// Hashing defines hashing of the data and prev block
func (b *Block) Hashing() {
	info := bytes.Join([][]byte{b.Data, b.Prev}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
