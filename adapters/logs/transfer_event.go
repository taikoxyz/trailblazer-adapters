package logs

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/erc20"
)

var (
	logTransferSigHash = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
)

// TransferIndexer is an implementation of LogsIndexer for ERC20 transfer logs.
type TransferIndexer struct {
	TargetAddresses []common.Address
}

// NewTransferIndexer creates a new TransferIndexer.
func NewTransferIndexer() *TransferIndexer {
	return &TransferIndexer{TargetAddresses: nil}
}

func (indexer *TransferIndexer) Addresses() []common.Address {
	return indexer.TargetAddresses
}

// IndexLogs processes logs for ERC20 transfers.
func (indexer *TransferIndexer) IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]adapters.Whitelist, error) {
	var result []adapters.Whitelist
	for _, vLog := range logs {
		if !isERC20Transfer(vLog) {
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

func isERC20Transfer(vLog types.Log) bool {
	return len(vLog.Topics) == 3 && vLog.Topics[0].Hex() == logTransferSigHash.Hex()
}

// processLog processes a single ERC20 transfer log.
func (indexer *TransferIndexer) ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*adapters.Whitelist, error) {
	to := common.BytesToAddress(vLog.Topics[2].Bytes()[12:])

	var transferEvent struct {
		Value *big.Int
	}

	erc20ABI, err := abi.JSON(strings.NewReader(erc20.ABI))
	if err != nil {
		return nil, err
	}
	err = erc20ABI.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
	if err != nil {
		return nil, err
	}

	block, err := client.BlockByNumber(ctx, big.NewInt(int64(vLog.BlockNumber)))
	if err != nil {
		return nil, err
	}

	return &adapters.Whitelist{
		User: to,
		Time: block.Time(),
	}, nil
}
