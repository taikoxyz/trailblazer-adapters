package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/okx"
)

func processOrderFulfilledIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := okx.NewOrderFulfilledIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
