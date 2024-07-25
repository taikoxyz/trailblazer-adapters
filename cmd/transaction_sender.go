package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/blocks"
)

func processNewTransactionSender(client *ethclient.Client, blockNumber int64) {
	processor := blocks.NewTransactionSender()
	processTransactionIndexer(client, processor, blockNumber)
}

func processNewNftDeployed(client *ethclient.Client, blockNumber int64) {
	processor := blocks.NewNftDeployedIndexer()
	processTransactionIndexer(client, processor, blockNumber)
}
