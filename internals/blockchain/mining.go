package blockchain

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ktruedat/goBlockchain/utils"
	"log"
	"strings"
)

const (
	MiningDifficulty = 3
	MiningSender     = "THE_BLOCKCHAIN"
	MiningReward     = 1.0
)

// Mining section
func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MiningSender, bc.blockchainAddress, MiningReward, nil, nil)
	nonce := bc.proofOfWork()
	previousHash := bc.LastBlock().Hash()
	bc.CreateBlock(nonce, previousHash)
	log.Println("mining completed successfully")
	return true
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32, senderPublicKey *ecdsa.PublicKey, s *utils.Signature) error {
	t := NewTransaction(sender, recipient, value)
	if sender == MiningSender {
		bc.transactionPool = append(bc.transactionPool, t)
		return nil
	}
	if bc.verifyTransactionSignature(senderPublicKey, s, t) {
		if bc.calculateTotalAmount(sender) < value {
			return errors.New("failed to send cryptocurrency: not enough balance in wallet")
		}
		bc.transactionPool = append(bc.transactionPool, t)
		return nil
	}
	return errors.New("verify transaction signature unsuccessful")
}

func (bc *Blockchain) verifyTransactionSignature(senderPublicKey *ecdsa.PublicKey, s *utils.Signature, t *Transaction) bool {
	m, _ := json.Marshal(t)
	h := sha256.Sum256(m[:])
	return ecdsa.Verify(senderPublicKey, h[:], s.R, s.S)
}

func (bc *Blockchain) calculateTotalAmount(blockchainAddr string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			if blockchainAddr == t.recipientBlockchainAddress {
				totalAmount += value
			}
			if blockchainAddr == t.senderBlockchainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

func (bc *Blockchain) copyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, t)
	}
	return transactions
}

func (bc *Blockchain) proofOfWork() int {
	transactions := bc.copyTransactionPool()
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
