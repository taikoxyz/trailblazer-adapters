package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/loopex"
)

func processNewSaleIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := loopex.NewNewSaleIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
