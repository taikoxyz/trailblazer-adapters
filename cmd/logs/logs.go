package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func main() {
	client, err := ethclient.Dial("https://rpc.taiko.xyz")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	processor := logs.NewDotTaikoIndexer([]common.Address{common.HexToAddress("0xD7b837A0E388B4c25200983bdAa3EF3A83ca86b7"), common.HexToAddress("0xFb2Cd41a8aeC89EFBb19575C6c48d872cE97A0A5")})

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch the block: %v", err)
	}
	query := ethereum.FilterQuery{
		Addresses: processor.Addresses,
		FromBlock: big.NewInt(int64(178747)),
		ToBlock:   big.NewInt(int64(178798)),
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
