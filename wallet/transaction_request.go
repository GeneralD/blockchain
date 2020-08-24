package wallet

type TransactionRequest struct {
	SenderPrivateKey           *string `json:"sender_private_key"`
	SenderPublicKey            *string `json:"sender_public_key"`
	SenderBlockchainAddress    *string `json:"sender_blockchain_address"`
	RecipientBlockchainAddress *string `json:"recipient_blockchain_address"`
	Value                      *string `json:"value"`
}

func (t *TransactionRequest) Validate() bool {
	return t.SenderPrivateKey != nil &&
		t.SenderPublicKey != nil &&
		t.SenderBlockchainAddress != nil &&
		t.RecipientBlockchainAddress != nil &&
		t.Value != nil
}
