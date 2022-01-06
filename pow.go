package blockchain

import (
    "fmt"
    "bytes"
    "math"
    "math/big"
    "crypto/sha256"
)

const difficulty = 12

// ProofOfWork defines structure for proof of work
type ProofOfWork struct {
	Block *Block
	Target *big.Int
}

// NewProofOfWork creates initialization of proof of work
func NewProofOfWork(b *Block) *ProofOfWork {
    target := big.NewInt(1)
    target.Lsh(target, uint(256-difficulty))
    pow := &ProofOfWork{b, target}
    return pow
}

func (pow *ProofOfWork) MakeNonce(nonce int64) ([]byte, error) {
    nonceHex, err := ToHex(nonce)
    if err != nil {
        return nil, err
    }
    difficultyHash, err := ToHex(difficulty)
    if err != nil {
        return nil, err
    }
    data := bytes.Join(
        [][]byte{
            pow.Block.Prev,
            pow.Block.Data,
            nonceHex,
            difficultyHash,
        },
        []byte{},
    )
    return data, nil
}

func (pow *ProofOfWork) Validate() bool {
    var intHash big.Int

    data, err := pow.MakeNonce(int64(pow.Block.Nonce))
    if err != nil {
        return false
    }

    hash := sha256.Sum256(data)
    intHash.SetBytes(hash[:])

    return intHash.Cmp(pow.Target) == -1
}

func (pow *ProofOfWork) Run() (int, []byte) {
    var intHash big.Int
    var hash [32]byte

    nonce := 0
    for nonce < math.MaxInt64 {
        data, err := pow.MakeNonce(int64(nonce))
        if err != nil {
            continue
        }
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