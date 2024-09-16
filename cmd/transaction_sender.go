package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/blocks"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/gaming"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/swing"
)

func processNewTransactionSender(client *ethclient.Client, blockNumber int64) error {
	processor := blocks.NewTransactionSender()
	return processTransactionIndexer(client, processor, blockNumber)
}

func processNewNftDeployed(client *ethclient.Client, blockNumber int64) error {
	processor := blocks.NewNftDeployedIndexer()
	return processTransactionIndexer(client, processor, blockNumber)
}

func processNewGamingWhitelist(client *ethclient.Client, blockNumber int64) error {
	processor := gaming.NewGamingWhitelist()
	return processTransactionIndexer(client, processor, blockNumber)
}

func processNewBridgeIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := swing.NewCrosschainSwapRequestIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
