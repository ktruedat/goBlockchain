package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func main() {
	//blockChain := NewBlockchain()
	//blockChain.CreateBlock(5, "hash 1")
	//blockChain.CreateBlock(3, "hash 2")
	//blockChain.Print()
	block := NewBlock(0, "init hash")
	hash := block.Hash()
	fmt.Printf("%x", hash)

}

type Block struct {
	nonce        int
	previousHash string
	timpestamp   int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timpestamp:   time.Now().UnixNano(),
		transactions: nil,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp 		%d\n", b.timpestamp)
	fmt.Printf("previousHash 		%s\n", b.previousHash)
	fmt.Printf("nonce 			%d\n", b.nonce)
	fmt.Printf("transactions 		%s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	fmt.Println(string(m))
	return sha256.Sum256(m)
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64    `json:"timestamp"`
		Nonce        int      `json:"nonce"`
		PreviousHash string   `json:"previousHash"`
		Transactions []string `json:"transactions"`
	}{Timestamp: b.timpestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions})
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) Print() {
	fmt.Printf("%s\n", strings.Repeat("*", 25))
	for i, block := range bc.chain {
		fmt.Printf("%s chain %d %s\n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}
