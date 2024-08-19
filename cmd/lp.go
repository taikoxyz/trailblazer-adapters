package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/izumi"
	"github.com/taikoxyz/trailblazer-adapters/adapters/projects/ritsu"
)

func processRitsuLPIndexer(client *ethclient.Client, blockNumber int64) error {
	whitelist := map[string]struct{}{
		common.HexToAddress("0xaE7850cBbA0d7303eD06661c9B06f4A5127Ca75D").Hex(): {},
		common.HexToAddress("0xd7506E9Ebac0Ff14Ce3C600DC6Fd94240e5612D4").Hex(): {},
	}
	tokens := []common.Address{
		common.HexToAddress("0x7c38E9389B27668280E5aaAc372eBCb2ECc1c5E0"),
		common.HexToAddress("0x6c7839E0CE8AdA360a865E18a111A462d08DC15a"),
	}
	processor := ritsu.NewTransferIndexer(tokens, whitelist)
	return processLPLogIndexer(client, processor, blockNumber)
}

func processIziLPIndexer(client *ethclient.Client, blockNumber int64) error {
	whitelist := map[string]struct{}{
		common.HexToAddress("0x88867BF3bB3321d8c7Da71a8eAb70680037068E4").Hex(): {},
	}
	tokens := []common.Address{
		common.HexToAddress("0x33531bDBFE34fa6Fd5963D0423f7699775AacaaF"),
	}
	processor := izumi.NewTransferIndexer(tokens, whitelist)
	return processLPLogIndexer(client, processor, blockNumber)
}
