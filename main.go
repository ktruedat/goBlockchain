package main

import (
	"fmt"
	"github.com/ktruedat/goBlockchain/wallet"
)

func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKeyStr())
	fmt.Println()
	fmt.Println(w.PublicKeyStr())
}
