package blockchain

import "bytes"

// IsBlockValid checks validation of prev and current blocks
// on blockchain
func IsBlockValid(prevBlock, block *Block) bool {
	if bytes.Compare(block.Prev, prevBlock.Hash) == 0 {
		return false
	}
	if block.ID != prevBlock.ID+1 {
		return false
	}
	return true
}