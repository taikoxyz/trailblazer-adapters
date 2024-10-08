package dzap

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/dzap"
)

var (
	logMultiSwappedSigHash       = crypto.Keccak256Hash([]byte("MultiSwapped(bytes32,address,address,address,tuple(address,address,address,uint256,uint256,uint256)[])"))
	logSwappedSigHash            = crypto.Keccak256Hash([]byte("Swapped(bytes32,address,address,address,tuple(address,address,address,uint256,uint256,uint256))"))
	logSwappedSingleTokenSigHash = crypto.Keccak256Hash([]byte("SwappedSingleToken(bytes32,address,address,tuple(address,address,address,uint256,uint256,uint256)"))
)

type SwapCompletedIndexer struct {
	TargetAddresses []common.Address
}

type SwapDetail struct {
	Dex                common.Address
	FromToken          common.Address
	ToToken            common.Address
	FromAmount         *big.Int
	LeftOverFromAmount *big.Int
	ReturnToAmount     *big.Int
}

type MultiSwappedEvent struct {
	TransactionId [32]byte
	Integrator    common.Address
	Sender        common.Address
	Recipient     common.Address
	SwapInfo      []SwapDetail
}

type SwappedEvent struct {
	TransactionId [32]byte
	Integrator    common.Address
	Sender        common.Address
	Recipient     common.Address
	SwapInfo      SwapDetail
}

type SwappedSingleTokenEvent struct {
	TransactionId [32]byte
	Sender        common.Address
	Recipient     common.Address
	SwapInfo      SwapDetail
}

func NewSwapCompletedIndexer() *SwapCompletedIndexer {
	return &SwapCompletedIndexer{TargetAddresses: []common.Address{
		common.HexToAddress("0xF708e11A7C94abdE8f6217B13e6fE39C8b9cC0a6"),
	}}
}

func (indexer *SwapCompletedIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *SwapCompletedIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
	var result []adapters.Whitelist
	for _, vLog := range logs {
		if indexer.isRelevantMultiSwappedLog(vLog.Topics[0]) {
			multiswappedData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
			if err != nil {
				return nil, err
			}
			result = append(result, *multiswappedData)
		} else if indexer.isRelevantSingleTokenSwappedLog(vLog.Topics[0]) {
			singleswappedData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
			if err != nil {
				return nil, err
			}
			result = append(result, *singleswappedData)

		} else if indexer.isRelevantSwappedLog(vLog.Topics[0]) {
			swappedData, err := indexer.ProcessLog(ctx, chainID, client, vLog)
			if err != nil {
				return nil, err
			}
			result = append(result, *swappedData)
		} else {
			continue
		}
	}
	return result, nil
}

func (indexer *SwapCompletedIndexer) isRelevantSwappedLog(topic common.Hash) bool {
	return topic.Hex() == logSwappedSigHash.Hex()
}
func (indexer *SwapCompletedIndexer) isRelevantMultiSwappedLog(topic common.Hash) bool {
	return topic.Hex() == logMultiSwappedSigHash.Hex()
}
func (indexer *SwapCompletedIndexer) isRelevantSingleTokenSwappedLog(topic common.Hash) bool {
	return topic.Hex() == logSwappedSingleTokenSigHash.Hex()
}

func (indexer *SwapCompletedIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
	var swappedSingleTokenEvent SwappedSingleTokenEvent
	var swappedEvent SwappedEvent
	var multiswappedEvent MultiSwappedEvent

	swapABI, err := abi.JSON(strings.NewReader(dzap.ABI))
	if err != nil {
		return nil, err
	}
	errSingleSwapped := swapABI.UnpackIntoInterface(&swappedSingleTokenEvent, "SwappedSingleToken", vLog.Data)
	errSwapped := swapABI.UnpackIntoInterface(&swappedEvent, "Swapped", vLog.Data)
	errMultiSwapped := swapABI.UnpackIntoInterface(&multiswappedEvent, "MultiSwapped", vLog.Data)

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	if errSingleSwapped == nil {
		return &adapters.Whitelist{
			User:        swappedSingleTokenEvent.Recipient,
			Time:        block.Time(),
			BlockNumber: block.Number().Uint64(),
		}, nil
	} else if errSwapped == nil {
		return &adapters.Whitelist{
			User:        swappedEvent.Recipient,
			Time:        block.Time(),
			BlockNumber: block.Number().Uint64(),
		}, nil
	} else if errMultiSwapped == nil {
		return &adapters.Whitelist{
			User:        multiswappedEvent.Recipient,
			Time:        block.Time(),
			BlockNumber: block.Number().Uint64(),
		}, nil
	}
	return nil, nil
}
