package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	SenderBlockchainAddress    string
	RecipientBlockchainAddress string
	Value                      float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{
		SenderBlockchainAddress:    sender,
		RecipientBlockchainAddress: recipient,
		Value:                      value,
	}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("sender_blockchain_addrees:		%s\n", t.SenderBlockchainAddress)
	fmt.Printf("recipient_blockchain_addrees:		%s\n", t.RecipientBlockchainAddress)
	fmt.Printf("Value:					%.1f\n", t.Value)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"SenderBlockchainAddress"`
		RecipientBlockchainAddress string  `json:"RecipientBlockchainAddress"`
		Value                      float32 `json:"Value"`
	}{SenderBlockchainAddress: t.SenderBlockchainAddress,
		RecipientBlockchainAddress: t.RecipientBlockchainAddress,
		Value:                      t.Value})
}
