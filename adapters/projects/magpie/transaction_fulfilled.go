package magpie

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/magpie"
)

var (
	logTransactionSigHash = crypto.Keccak256Hash([]byte("Swap(address,address,address,address,uint256,uint256)"))
)

type TransactionIndexer struct {
	TargetAddresses []common.Address
}

type TransactionEvent struct {
	FromAddress      common.Address `json:"fromAddress"`
	ToAddress        common.Address `json:"toAddress"`
	FromAssetAddress common.Address `json:"fromAssetAddress"`
	ToAssetAddress   common.Address `json:"toAssetAddress"`
	AmountIn         *big.Int       `json:"amountIn"`
	AmountOut        *big.Int       `json:"amountOut"`
}

func MagpieTransactionFulfilledIndexer() *TransactionIndexer {
	return &TransactionIndexer{TargetAddresses: []common.Address{
		common.HexToAddress("0x956df8424b556f0076e8abf5481605f5a791cc7f"),
	}}
}

func (indexer *TransactionIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *TransactionIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
	var result []adapters.Whitelist
	// fmt.Println(logs)
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

func (indexer *TransactionIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == logTransactionSigHash.Hex()
}

func (indexer *TransactionIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
	var TransactionEvent TransactionEvent
	TransactionABI, err := abi.JSON(strings.NewReader(magpie.ABI))
	if err != nil {
		return nil, err
	}
	err = TransactionABI.UnpackIntoInterface(&TransactionEvent, "Swap", vLog.Data)
	if err != nil {
		return nil, err
	}

	var sender = vLog.Topics[1]
	address := common.BytesToAddress(sender.Bytes()[12:])
	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}
	return &adapters.Whitelist{
		User:        address,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
