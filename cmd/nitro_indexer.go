package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/nitro"
)

func processNitroIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := nitro.NewIRelayIndexer()
	return processLogIndexer(client, processor, blockNumber)
}