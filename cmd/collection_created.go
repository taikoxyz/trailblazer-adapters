package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/nfts2me"
)

func processCollectionCreatedIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := nfts2me.NewCollectionCreatedIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
