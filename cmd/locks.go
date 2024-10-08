package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/drips"
)

const dripsLockContract = "0x46f0a2e45bee8e9ebfdb278ce06caa6af294c349"

func processDripsIndexer(client *ethclient.Client, blockNumber int64) error {
	processor := drips.NewDripsIndexer(common.HexToAddress(dripsLockContract), client)
	return processLockIndexer(client, processor, blockNumber)
}
