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
	Addresses() []common.Address
	IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]Whitelist, error)
	ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*Whitelist, error)
}

type LPLogsIndexer interface {
	Address() []common.Address
	IndexLogs(ctx context.Context, chainID *big.Int, client *ethclient.Client, logs []types.Log) ([]LPTransfer, error)
	ProcessLog(ctx context.Context, chainID *big.Int, client *ethclient.Client, vLog types.Log) (*LPTransfer, error)
}

type LockLogsIndexer interface {
	Address() []common.Address
	IndexLogs(ctx context.Context, logs []types.Log) ([]Lock, error)
	ProcessLog(ctx context.Context, vLog types.Log) (*Lock, error)
}
