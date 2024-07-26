package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processNewSaleIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := logs.NewNewSaleIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
