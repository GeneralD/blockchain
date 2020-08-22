package blockchain

import (
	"encoding/json"
	"strings"
)

type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBlockchainAddress    string  `json:"sender_blockchain_address"`
		RecipientBlockchainAddress string  `json:"recipient_blockchain_address"`
		Value                      float32 `json:"value"`
	}{
		SenderBlockchainAddress:    t.senderBlockchainAddress,
		RecipientBlockchainAddress: t.recipientBlockchainAddress,
		Value:                      t.value,
	})
}

func (t *Transaction) Print() {
	Logger.Tracef(strings.Repeat("- ", 40))
	Logger.Tracef("sender_blockchain_address      %s", t.senderBlockchainAddress)
	Logger.Tracef("recipient_blockchain_address   %s", t.recipientBlockchainAddress)
	Logger.Tracef("value                          %.1f", t.value)
}
