package blockchain

import "errors"

type BlockChain struct {
	Blocks []*Block
}

// AddBlock provides adding of the new block to blockchain
func (chain *BlockChain) AddBlock(data []byte) error  {
    prevBlock := chain.Blocks[len(chain.Blocks)-1]
    newBlock := CreateBlock(data, prevBlock.Hash)
    if ok := IsBlockValid(prevBlock, newBlock); ok {
        return errors.New("unable to validate blocks")
    }
    chain.Blocks = append(chain.Blocks, newBlock)
    return nil
}

// Create provides creating of the new block with genesis
func Create()*BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}

// ChooseBlockChain provides select better blockchain between 
// remote and local
func ChooseBlockChain(local, remote *BlockChain) (*BlockChain, error) {
    isRemoteValid := IsChainValid(remote)
    isLocalValid := IsChainValid(local)
    if isRemoteValid && isLocalValid {
        if len(remote.Blocks) > len(local.Blocks) {
            return remote, nil
        }
        return local, nil
    }
    if isRemoteValid && !isLocalValid {
        return remote, nil
    }
    if !isRemoteValid && isLocalValid {
        return local, nil
    }
    return nil, errors.New("local and remote blockchains is invalid")

}