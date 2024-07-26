package cmd

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

func processLogIndexer(client *ethclient.Client, processor adapters.TransferLogsIndexer, blockNumber int64) error {
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch the chain ID: %v", err)
		return err
	}
	query := ethereum.FilterQuery{
		Addresses: processor.Addresses(),
		FromBlock: big.NewInt(blockNumber),
		ToBlock:   big.NewInt(blockNumber),
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("Failed to fetch the logs: %v", err)
		return err
	}
	senders, err := processor.IndexLogs(context.Background(), chainID, client, logs)
	if err != nil {
		log.Fatalf("Failed to process the logs: %v", err)
		return err
	}

	fmt.Printf("Senders: %v\n", senders)
	return nil
}
