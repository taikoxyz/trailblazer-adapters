package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processOrderFulfilledIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := logs.NewOrderFulfilledIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
