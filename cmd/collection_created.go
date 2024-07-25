package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processCollectionCreatedIndexer(client *ethclient.Client, blockNumber int64) {
	processor := logs.NewCollectionCreatedIndexer()
	processLogIndexer(client, processor, blockNumber)
}
