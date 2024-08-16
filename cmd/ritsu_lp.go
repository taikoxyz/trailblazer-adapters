package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/ritsu"
)

func processRitsuLPIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := ritsu.NewTransferIndexer(common.HexToAddress("0x7c38E9389B27668280E5aaAc372eBCb2ECc1c5E0"))
	return processLPLogIndexer(client, processor, blockNumber)
}
