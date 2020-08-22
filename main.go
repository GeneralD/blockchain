package main

import (
	"blockchain/blockchain"
	"blockchain/util"
	"blockchain/wallet"
)

func main() {
	bc := blockchain.NewBlockchain("miner_blockchain_address")
	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()
	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.0)
	bc.Mining()
	bc.Print()

	util.Logger.Debugf("miner_blockchain_address %.1f\n", bc.TotalAmount("miner_blockchain_address"))
	util.Logger.Debugf("C %.1f\n", bc.TotalAmount("C"))
	util.Logger.Debugf("D %.1f\n", bc.TotalAmount("D"))

	w := wallet.NewWallet()
	util.Logger.Debugf("PrivateKey: %s", w.PrivateKeyString())
	util.Logger.Debugf("PublicKey:  %s", w.PublicKeyString())
	util.Logger.Debugf("Address:    %s", w.Address())

	t := wallet.NewTransaction(w.PrivateKey(), w.PublicKey(), w.Address(), "B", 1.0)
	util.Logger.Debugf("signature:  %s", t.GenerateSignature())
}
