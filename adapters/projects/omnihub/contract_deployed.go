package omnihub

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

var (
	logContractDeployedSigHash = crypto.Keccak256Hash([]byte("ContractDeployed(address)"))
)

type ContractDeployedIndexer struct {
	TargetAddresses []common.Address
}

func NewContractDeployedIndexer() *ContractDeployedIndexer {
	return &ContractDeployedIndexer{TargetAddresses: []common.Address{
		common.HexToAddress("0xb0B4B761C9e9Bf5A9194a42e944A4A6646B83919"),
	}}
}

func (indexer *ContractDeployedIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

func (indexer *ContractDeployedIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
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

func (indexer *ContractDeployedIndexer) isRelevantLog(topic common.Hash) bool {
	return topic.Hex() == logContractDeployedSigHash.Hex()
}

func (indexer *ContractDeployedIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
	txn, err := client.TransactionInBlock(ctx, vLog.BlockHash, vLog.TxIndex)
	if err != nil {
		return nil, err
	}

	sender, err := client.TransactionSender(ctx, txn, vLog.BlockHash, vLog.TxIndex)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.Whitelist{
		User:        sender,
		Time:        block.Time(),
		BlockNumber: block.Number().Uint64(),
	}, nil
}
