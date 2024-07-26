package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/blocks"
)

func processNewTransactionSender(client *ethclient.Client, blockNumber int64) error {
	processor := blocks.NewTransactionSender()
	return processTransactionIndexer(client, processor, blockNumber)
}

func processNewNftDeployed(client *ethclient.Client, blockNumber int64) error {
	processor := blocks.NewNftDeployedIndexer()
	return processTransactionIndexer(client, processor, blockNumber)
}
