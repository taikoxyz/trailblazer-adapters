package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/conft"
)

func processTokenSoldIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := conft.NewTokenSoldIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
