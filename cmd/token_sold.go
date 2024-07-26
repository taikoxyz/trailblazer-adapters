package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processTokenSoldIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := logs.NewTokenSoldIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
