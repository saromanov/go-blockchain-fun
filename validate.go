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

// IsChainValid checks is blockchain is valid
func IsChainValid(chain *BlockChain) bool {
	for i := 1; i < len(chain.Blocks); i++ {
		firstBlock := chain.Blocks[i-1]
		secondBlock := chain.Blocks[i]
		if !IsBlockValid(secondBlock, firstBlock) {
			return false
		}
	}
	return true
}