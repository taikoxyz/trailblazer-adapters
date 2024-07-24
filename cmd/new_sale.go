package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processNewSaleIndexer(client *ethclient.Client, blockNumber int64) {
	processor := logs.NewNewSaleIndexer()
	processLogIndexer(client, processor, blockNumber)
}
