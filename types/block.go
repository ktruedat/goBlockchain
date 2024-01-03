package types

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timpestamp   int64
	Transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timpestamp:   time.Now().UnixNano(),
		Transactions: transactions,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp 		%d\n", b.timpestamp)
	fmt.Printf("nonce: 			%d\n", b.nonce)
	fmt.Printf("previousHash: 		%x\n", b.previousHash)
	for _, t := range b.Transactions {
		t.Print()
	}

}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"`
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previousHash"`
		Transactions []*Transaction `json:"Transactions"`
	}{Timestamp: b.timpestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.Transactions})
}
