package cmd

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/dzap"
)

func processDzapTransaction(client *ethclient.Client, blockNumber int64) error {
	processor := dzap.NewSwapCompletedIndexer()
	return processLogIndexer(client, processor, blockNumber)
}
