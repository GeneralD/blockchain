package blockchain

import (
	"blockchain/util"
	"fmt"
	"strings"
)

const (
	MiningDifficulty = 3
	MiningSender     = "THE BLOCKCHAIN"
	MiningReward     = 1.0
)

type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

// Instantiate a blockchain
func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{} // initial block
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.chainNewBlock(0, b.Hash())
	return bc
}

// Add a new transaction in pool
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

// Add transaction to earn reward and work to create a new block.
func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MiningSender, bc.blockchainAddress, MiningReward)
	bc.createBlock()
	util.Logger.Tracef("action=mining, status=success")
	return true
}

// Calculate total amount of the address
func (bc *Blockchain) TotalAmount(blockchainAddress string) float32 {
	var amount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			if blockchainAddress == t.recipientBlockchainAddress {
				amount += t.value
			}
			if blockchainAddress == t.senderBlockchainAddress {
				amount -= t.value
			}
		}
	}
	return amount
}

// Work to create a new block and chain it to blockchain
func (bc *Blockchain) createBlock() {
	nonce := bc.proofOfWork()
	previousHash := bc.lastBlock().Hash()
	bc.chainNewBlock(nonce, previousHash)
}

// Create a new block and chain it to blockchain
func (bc *Blockchain) chainNewBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

// Last block in the blockchain
func (bc *Blockchain) lastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

// Hard copy current transaction pool
func (bc *Blockchain) copyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value))
	}
	return transactions
}

// Check if the nonce is valid
func (bc *Blockchain) validProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	b := Block{nonce, previousHash, 0, transactions}
	hashStr := fmt.Sprintf("%x", b.Hash())
	return hashStr[:difficulty] == strings.Repeat("0", difficulty)
}

// Find nice nonce to pass the validation
func (bc *Blockchain) proofOfWork() int {
	transactions := bc.copyTransactionPool()
	previousHash := bc.lastBlock().Hash()
	nonce := 0
	for !bc.validProof(nonce, previousHash, transactions, MiningDifficulty) {
		nonce++
	}
	return nonce
}

// Print the blockchain
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		util.Logger.Tracef("%s Block %d %s", strings.Repeat("=", 35), i, strings.Repeat("=", 35))
		block.Print()
	}
	util.Logger.Tracef(strings.Repeat("*", 80))
}
