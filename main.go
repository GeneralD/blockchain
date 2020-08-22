package main

import (
	"blockchain/blockchain"
	"blockchain/utils"
	"blockchain/wallet"
)

func main() {
	minerWallet := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet side
	transaction := walletA.SendTo(walletB.Address(), 1.0)
	signature := transaction.GenerateSignature()

	// Blockchain side
	blockchain := blockchain.NewBlockchain(minerWallet.Address())
	isAdded := blockchain.AddTransaction(walletA.Address(), walletB.Address(), 1.0, walletA.PublicKey(), signature)
	utils.Logger.Debugf("Result: %t", isAdded)

	blockchain.Mining()
	blockchain.Print()

	utils.Logger.Infof("A     %.1f", blockchain.TotalAmount(walletA.Address()))
	utils.Logger.Infof("B     %.1f", blockchain.TotalAmount(walletB.Address()))
	utils.Logger.Infof("Miner %.1f", blockchain.TotalAmount(minerWallet.Address()))
}
