package blockchain

import (
	"blockchain/utils"
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/json"
	"strings"
)

type Transaction struct {
	senderAddress    string
	recipientAddress string
	value            float32
}

func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

func (t *Transaction) VerifySignature(senderPublicKey *ecdsa.PublicKey, signature *utils.Signature) bool {
	m, _ := json.Marshal(t)
	hash := sha256.Sum256(m)
	return ecdsa.Verify(senderPublicKey, hash[:], signature.R, signature.S)
}

func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderAddress    string  `json:"sender_address"`
		RecipientAddress string  `json:"recipient_address"`
		Value            float32 `json:"value"`
	}{
		SenderAddress:    t.senderAddress,
		RecipientAddress: t.recipientAddress,
		Value:            t.value,
	})
}

func (t *Transaction) Print() {
	utils.Logger.Tracef(strings.Repeat("- ", 40))
	utils.Logger.Tracef("sender_blockchain_address      %s", t.senderAddress)
	utils.Logger.Tracef("recipient_blockchain_address   %s", t.recipientAddress)
	utils.Logger.Tracef("value                          %.1f", t.value)
}
