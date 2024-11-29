package main

import "fmt"

func main() {
	wallet := NewWallet()
	fmt.Printf("Wallet address: %s\n", wallet.GetAddress())
	bc := CreateBlockchain(string(wallet.GetAddress()))
	defer bc.db.Close()
	bc.getBalance(string(wallet.GetAddress()))
}
