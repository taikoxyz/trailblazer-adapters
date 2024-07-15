package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/blocks"
)

func main() {
	client, err := ethclient.Dial("https://rpc.taiko.xyz")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	validSenders := map[string]struct{}{
		// testing for transaction https://taikoscan.io/tx/0x60d3a713b296055a4dd8f5e2a0e9094ffc5aaf0270a61c3c037d98f4d261c3be
		common.HexToAddress("0x1Df2De291F909baA50C1456C87C71Edf9Fb199D5").Hex(): {},
	}
	processor := blocks.NewTransactionSender(validSenders)
	blockNumber := big.NewInt(1000) // Replace with the block number you want to process
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatalf("Failed to fetch the block: %v", err)
	}

	senders, err := processor.ProcessBlock(context.Background(), block, client)
	if err != nil {
		log.Fatalf("Failed to process the block: %v", err)
	}

	fmt.Printf("Senders: %v\n", senders)
}
