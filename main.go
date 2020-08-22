package main

import (
	"blockchain/blockchain"
	"blockchain/utils"
	"blockchain/wallet"
)

func main() {
	//bc := blockchain.NewBlockchain("miner_blockchain_address")
	//bc.AddTransaction("A", "B", 1.0)
	//bc.Mining()
	//bc.AddTransaction("C", "D", 2.0)
	//bc.AddTransaction("X", "Y", 3.0)
	//bc.Mining()
	//bc.Print()
	//
	//utils.Logger.Debugf("miner_blockchain_address %.1f\n", bc.TotalAmount("miner_blockchain_address"))
	//utils.Logger.Debugf("C %.1f\n", bc.TotalAmount("C"))
	//utils.Logger.Debugf("D %.1f\n", bc.TotalAmount("D"))
	//
	//w := wallet.NewWallet()
	//utils.Logger.Debugf("PrivateKey: %s", w.PrivateKeyString())
	//utils.Logger.Debugf("PublicKey:  %s", w.PublicKeyString())
	//utils.Logger.Debugf("Address:    %s", w.Address())
	//
	//t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.Address(), "B", 1.0)
	//utils.Logger.Debugf("signature:  %s", t.GenerateSignature())

	//walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	// Wallet side
	transaction := walletA.SendTo(walletB.Address(), 1.0)
	signature := transaction.GenerateSignature()

	// Blockchain side
	blockchain := blockchain.NewBlockchain("miner_blockchain_address")
	isAdded := blockchain.AddTransaction(walletA.Address(), walletB.Address(), 1.0, walletA.PublicKey(), signature)
	utils.Logger.Debugf("Result: %t", isAdded)
}
