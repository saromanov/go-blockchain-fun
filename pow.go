package blockchain

import (
    "bytes"
    "math"
    "math/big"
)
// ProofOfWork defines structure for proof of work
type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

// NewProofOfWork creates initialization of proof of work
func NewProofOfWork(b *Block) *ProofOfWork {
    target := big.NewInt(1)
    target.Lsh(target, uint(256-Difficulty))
    pow := &ProofOfWork{b, target}
    return pow
}

func (pow *ProofOfWork) MakeNonce(nonce, difficylty int64) []byte {
    data := bytes.Join(
        [][]byte{
            pow.Block.PrevHash,
            pow.Block.Data,
            ToHex(nonce),
            ToHex(difficulty),
        },
        []byte{},
    )
    return data
}

func (pow *ProofOfWork) Run() (int, []byte) {
    var intHash big.Int
    var hash [32]byte

    nonce := 0
    for nonce < math.MaxInt64 {
        data := pow.InitNonce(nonce)
        hash = sha256.Sum256(data)

        fmt.Printf("\r%x", hash)
        intHash.SetBytes(hash[:])

        if intHash.Cmp(pow.Target) == -1 {
            break
        } else {
            nonce++
        }

    }
    fmt.Println()

    return nonce, hash[:]
}