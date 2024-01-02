package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()

	previousHash := blockChain.LastBlock().Hash()
	blockChain.CreateBlock(5, previousHash)
	blockChain.Print()

	previousHash = blockChain.LastBlock().Hash()
	blockChain.CreateBlock(3, previousHash)
	blockChain.Print()

	//block := NewBlock(0, "init hash")
	//hash := block.Hash()
	//fmt.Printf("%x", hash)

}

type Block struct {
	nonce        int
	previousHash [32]byte
	timpestamp   int64
	transactions []string
}

func NewBlock(nonce int, previousHash [32]byte) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timpestamp:   time.Now().UnixNano(),
		transactions: nil,
	}
}

func (b *Block) Print() {
	fmt.Printf("timestamp 		%d\n", b.timpestamp)
	fmt.Printf("previousHash 		%x\n", b.previousHash)
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
		PreviousHash [32]byte `json:"previousHash"`
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
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
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

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}
