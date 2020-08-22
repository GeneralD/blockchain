package main

import (
	"blockchain/blockchain"
	"fmt"
)

func main() {
	bc := blockchain.NewBlockchain("miner_blockchain_address")
	bc.AddTransaction("A", "B", 1.0)
	bc.Mining()
	bc.AddTransaction("C", "D", 2.0)
	bc.AddTransaction("X", "Y", 3.0)
	bc.Mining()
	bc.Print()

	fmt.Printf("miner_blockchain_address %.1f\n", bc.TotalAmount("miner_blockchain_address"))
	fmt.Printf("C %.1f\n", bc.TotalAmount("C"))
	fmt.Printf("D %.1f\n", bc.TotalAmount("D"))
}
