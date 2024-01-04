package wallet

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
)

type Transaction struct {
	senderPrivateKey           *ecdsa.PrivateKey
	senderPublicKey            *ecdsa.PublicKey
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(senderPrivateKey *ecdsa.PrivateKey, senderPublicKey *ecdsa.PublicKey, senderBlockchainAddress string, recipientBlockchainAddress string, value float32) *Transaction {
	return &Transaction{senderPrivateKey: senderPrivateKey, senderPublicKey: senderPublicKey, senderBlockchainAddress: senderBlockchainAddress, recipientBlockchainAddress: recipientBlockchainAddress, value: value}
}

func (t *Transaction) GenerateSignature() *Signature {
	m, _ := json.Marshal(t)

}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"senderBlockchainAddress"`
		RecipientBlockchainAddress string  `json:"recipientBlockchainAddress"`
		Value                      float32 `json:"value"`
	}{SenderBlockchainAddress: t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value})
}

type Signature struct {
	R *big.Int
	S *big.Int
}
