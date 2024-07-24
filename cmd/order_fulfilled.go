package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processOrderFulfilledIndexer(client *ethclient.Client, blockNumber int64) {
	processor := logs.NewOrderFulfilledIndexer()
	processLogIndexer(client, processor, blockNumber)
}
