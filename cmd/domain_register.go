package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/domains"
)

func processDotTaikoIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := domains.NewDotTaikoIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
