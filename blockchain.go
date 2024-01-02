package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("test")
	b := NewBlock(0, "init hash")
	b.Print()
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

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string)
