package main

import (
	"fmt"
	"github.com/ktruedat/goBlockchain/types"
	"strings"
)

type Blockchain struct {
	transactionPool   []*types.Transaction
	chain             []*types.Block
	blockchainAddress string
}

func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &types.Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
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

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *types.Block {
	b := types.NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) LastBlock() *types.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) AddTransaction(sender, recipient string, value float32) {
	t := types.NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*types.Transaction {
	transactions := make([]*types.Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, t)
	}
	return transactions
}

func (bc *Blockchain) CalculateTotalAmount(blockchainAddr string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.Transactions {
			if blockchainAddr == t.
		}
	}
}
