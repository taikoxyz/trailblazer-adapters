package adapters

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// TransferLogsIndexer is an interface that defines the methods for indexing and processing transfer logs.
type TransferLogsIndexer interface {
	IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) (*[]TransferData, error)
	ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*TransferData, error)
}

// TransferData represents an event with specific fields.
type TransferData struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Time  uint64
}
