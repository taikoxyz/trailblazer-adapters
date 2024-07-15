package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func main() {
	client, err := ethclient.Dial("https://rpc.taiko.xyz")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	processor := logs.NewTransferIndexer(nil)

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch the block: %v", err)
	}
	query := ethereum.FilterQuery{
		Addresses: processor.Addresses,
		FromBlock: big.NewInt(int64(1000)),
		ToBlock:   big.NewInt(int64(1001)),
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to fetch the logs: %v", err)
	}
	senders, err := processor.IndexLogs(context.Background(), chainID, client, logs)
	if err != nil {
		log.Fatalf("Failed to process the block: %v", err)
	}

	fmt.Printf("Senders: %v\n", senders)
}
