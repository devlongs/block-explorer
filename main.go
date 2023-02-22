package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// Connect to the Ethereum mainnet
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_PROJECT_ID")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum mainnet: %v", err)
	}
	defer client.Close()

	// Prompt the user to enter a transaction hash or address
	fmt.Print("Enter a transaction hash or address: ")
	var input string
	fmt.Scanln(&input)

	// Check if the input is a valid transaction hash
	if _, err := common.NewHashFromHex(input); err == nil {
		// The input is a valid transaction hash, so display detailed information about the transaction
		transaction, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(input))
		if err != nil {
			log.Fatalf("Failed to retrieve transaction: %v", err)
		}
		fmt.Printf("Hash: %v\n", transaction.Hash().Hex())
		fmt.Printf("Nonce: %v\n", transaction.Nonce())
		fmt.Printf("Block Hash: %v\n", transaction.BlockHash().Hex())
		fmt.Printf("Block Number: %v\n", transaction.BlockNumber())
		fmt.Printf("Transaction Index: %v\n", transaction.Index())
		fmt.Printf("From: %v\n", transaction.From().Hex())
		fmt.Printf("To: %v\n", transaction.To().Hex())
		fmt.Printf("Value: %v\n", transaction.Value().String())
		fmt.Printf("Gas: %v\n", transaction.Gas())
		fmt.Printf("Gas Price: %v\n", transaction.GasPrice().String())
		fmt.Printf("Input: %v\n", transaction.Data())
		if isPending {
			fmt.Println("Status: Pending")
		} else {
			receipt, err := client.TransactionReceipt(context.Background(), transaction.Hash())
			if err != nil {
				log.Fatalf("Failed to retrieve receipt: %v", err)
			}
			if receipt.Status == types.ReceiptStatusSuccessful {
				fmt.Println("Status: Success")
			} else {
				fmt.Println("Status: Failed")
			}
		}
	} else {
		// The input is not a valid transaction hash, so assume it is an address and display information about the address
		address := common.HexToAddress(input)
		balance, err := client.BalanceAt(context.Background(), address, nil)
		if err != nil {
			log.Fatalf("Failed to retrieve balance: %v", err)
		}
		nonce, err := client.PendingNonceAt(context.Background(), address)
		if err != nil {
			log.Fatalf("Failed to retrieve nonce: %v", err)
		}
		fmt.Printf("Address: %v\n", address.Hex())
		fmt.Printf("Balance: %v\n", balance.String())
		fmt.Printf("Nonce: %v\n", nonce)
	}

	// Prompt the user to search for another transaction or address
	fmt.Print("Search again? (y/n): ")
	var response string
	fmt.Scanln(&response)
	if response == "y" {
		main()
	} else {
		os.Exit(0)
	}
}
