package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/eisen"
)

func processEisenIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := eisen.NewSwapsCompleteIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
