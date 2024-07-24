package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/logs"
)

func processDotTaikoIndexer(client *ethclient.Client, blockNumber int64) {
	processor := logs.NewDotTaikoIndexer()
	processLogIndexer(client, processor, blockNumber)
}
