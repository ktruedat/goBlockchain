package main

import (
	"github.com/ktruedat/goBlockchain/blockchain"
	"github.com/ktruedat/goBlockchain/wallet"
)

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	t := wallet.NewTransaction(walletA.PrivateKey(), walletB.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)
	bc := blockchain.NewBlockchain(walletM.BlockchainAddress())
	if err := bc.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature()); err != nil {
		panic(err)
	}
	//fmt.Println(w.PrivateKeyStr())
	//fmt.Println()
	//fmt.Println(w.PublicKeyStr())

}
