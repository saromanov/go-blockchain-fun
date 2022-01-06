package blockchain

import "bytes"

// IsValid checks validation of prev and current blocks
// on blockchain
func IsValid(prevBlock, block *Block) bool {
	if bytes.Compare(block.Prev, prevBlock.Hash) == 0 {
		return false
	}
	if block.ID != prevBlock.ID+1 {
		return false
	}
	return true
}