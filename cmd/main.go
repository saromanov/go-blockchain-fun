package main

import (
	"fmt"
	"strconv"
	
	blockchain "github.com/saromanov/go-blockchain-fun"

)

func main() {

    chain := blockchain.Create()

    chain.AddBlock([]byte("first block"))
    chain.AddBlock([]byte("second block"))
    chain.AddBlock([]byte("third block"))

    for _, block := range chain.Blocks {

        fmt.Printf("Previous hash: %x\n", block.Prev)
        fmt.Printf("data: %s\n", block.Data)
        fmt.Printf("hash: %x\n", block.Hash)

        pow := blockchain.NewProofOfWork(block)
        fmt.Printf("Pow: %s\n", strconv.FormatBool(pow.Validate()))
        fmt.Println()
    }

}