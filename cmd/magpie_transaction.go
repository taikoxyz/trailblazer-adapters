package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	omnihub "github.com/taikoxyz/trailblazer-adapters/adapters/projects/magpie"
)

func processMagpieTransactionFulfilledIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := omnihub.MagpieTransactionFulfilledIndexer()
	return processLogIndexer(client, processor, blockNumber)
}