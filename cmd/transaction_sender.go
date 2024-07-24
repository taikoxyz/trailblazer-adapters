package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/blocks"
)

func processNewTransactionSender(client *ethclient.Client) {
	processor := blocks.NewTransactionSender()
	blockNumberBig := big.NewInt(blockNumber)
	block, err := client.BlockByNumber(context.Background(), blockNumberBig)
	if err != nil {
		log.Fatalf("Failed to fetch the block: %v", err)
	}

	senders, err := processor.ProcessBlock(context.Background(), block, client)
	if err != nil {
		log.Fatalf("Failed to process the block: %v", err)
	}

	fmt.Printf("Senders: %v\n", senders)
}
