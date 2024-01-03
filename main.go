package main

import "github.com/ktruedat/goBlockchain/blockchain"

func main() {
	addr := "some address"
	blockChain := blockchain.NewBlockchain(addr)
	blockChain.Print()

	blockChain.AddTransaction("A", "B", 1.0)
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 2.0)
	blockChain.AddTransaction("X", "Y", 3.0)
	blockChain.Mining()
	blockChain.Print()

}
