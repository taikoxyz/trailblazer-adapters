package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processCollectionCreatedIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := logs.NewCollectionCreatedIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
