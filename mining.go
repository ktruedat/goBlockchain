package main

import (
	"fmt"
	"github.com/ktruedat/goBlockchain/types"
	"strings"
)

const MiningDifficulty = 3

func (bc *Blockchain) ValidProof(difficulty, nonce int, previousHash [32]byte, transactions []*types.Transaction) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := types.NewBlock(nonce, previousHash, transactions)
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(MiningDifficulty, nonce, previousHash, transactions) {
		nonce += 1
	}
	return nonce
}
