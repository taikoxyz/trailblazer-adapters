package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/chainspot"
)

func processChainspotIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := chainspot.NewChainspotIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
