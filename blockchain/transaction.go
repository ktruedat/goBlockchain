package blockchain

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender, recipient string, value float32) *Transaction {
	return &Transaction{
		senderBlockchainAddress:    sender,
		recipientBlockchainAddress: recipient,
		value:                      value,
	}
}

func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("sender_blockchain_addrees:		%s\n", t.senderBlockchainAddress)
	fmt.Printf("recipient_blockchain_addrees:		%s\n", t.recipientBlockchainAddress)
	fmt.Printf("value:					%.1f\n", t.value)
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
