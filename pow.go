package blockchain

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