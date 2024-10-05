package eisen

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/eisen"
)

var (
	logSwapCompletedSigHash = crypto.Keccak256Hash([]byte("EisenSwapCompleted(address,address,address,address,uint256,uint256,uint256,uint256)"))
)

type SwapsCompleteIndexer struct {
	TargetAddresses []common.Address
}


type EisenSwapCompletedEvent struct {
	OrderHash     [32]byte
	Sender       common.Address
	FromAssetId    common.Address
	ToAssetId      common.Address
	Receiver    common.Address
	FromAmount 	 *big.Int
	ToAmount     *big.Int
	ExpectedToAmount *big.Int
	Fee *big.Int
}

func NewSwapsCompleteIndexer() *SwapsCompleteIndexer {
	return &SwapsCompleteIndexer{TargetAddresses: []common.Address{
		common.HexToAddress("0xFA0e9251503DaE51670d10288e6962d63191731d"),
	}}
}

func (indexer *SwapsCompleteIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *SwapsCompleteIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
	var result []adapters.Whitelist
	for _, vLog := range logs {
		if !indexer.isRelevantLog(vLog.Topics[0]) {
			continue
		}
		transferData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
		if err != nil {
			return nil, err
		}
		result = append(result, *transferData)
	}
	return result, nil
}

func (indexer *SwapsCompleteIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == logSwapCompletedSigHash.Hex()
}

func (indexer *SwapsCompleteIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
	var swapCompleteEvent EisenSwapCompletedEvent

	swapExecutorABI, err := abi.JSON(strings.NewReader(eisen.ABI))
	if err != nil {
		return nil, err
	}
	err = swapExecutorABI.UnpackIntoInterface(&swapCompleteEvent, "EisenSwapCompleted", vLog.Data)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.Whitelist{
		User:        swapCompleteEvent.Receiver,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
