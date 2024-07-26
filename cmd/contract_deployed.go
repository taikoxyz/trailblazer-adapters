package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/omnihub"
)

func processContractDeployedIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := omnihub.NewContractDeployedIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
