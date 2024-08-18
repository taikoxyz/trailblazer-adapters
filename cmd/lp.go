package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/izumi"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/ritsu"
)

func processRitsuLPIndexer(client *ethclient.Client, blockNumber int64) error {
	whitelist := map[string]struct{}{
		common.HexToAddress("0x67Cab37067309B664CFB12769FffA2bdc41D8EDe").Hex(): {},
	}
	processor := ritsu.NewTransferIndexer(common.HexToAddress("0x7c38E9389B27668280E5aaAc372eBCb2ECc1c5E0"), whitelist)
	return processLPLogIndexer(client, processor, blockNumber)
}

func processIziLPIndexer(client *ethclient.Client, blockNumber int64) error {
	whitelist := map[string]struct{}{
		common.HexToAddress("0x67Cab37067309B664CFB12769FffA2bdc41D8EDe").Hex(): {},
	}
	processor := izumi.NewTransferIndexer(common.HexToAddress("0x33531bDBFE34fa6Fd5963D0423f7699775AacaaF"), whitelist)
	return processLPLogIndexer(client, processor, blockNumber)
}
