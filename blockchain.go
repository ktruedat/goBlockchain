package main

import (
	"fmt"
	"time"
)

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()

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

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "init hash")
	return bc
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("chain %d\n", i)
		block.Print()
	}
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}
