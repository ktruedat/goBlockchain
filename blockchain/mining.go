package blockchain

import (
	"fmt"
	"log"
	"strings"
)

const (
	MiningDifficulty = 3
	MINING_SENDER    = "THE_BLOCKCHAIN"
	MINING_REWARD    = 1.0
)

func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD, nil, nil)
	nonce := bc.proofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("mining completed successfully")
	return true
}

func (bc *Blockchain) proofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.validProof(MiningDifficulty, nonce, previousHash, transactions) {
		nonce += 1
	}
	return nonce
}

func (bc *Blockchain) validProof(difficulty, nonce int, previousHash [32]byte, transactions []*Transaction) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := NewBlock(nonce, previousHash, transactions)
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}
