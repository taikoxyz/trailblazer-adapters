package adapters

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
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

type Lock struct {
	User          common.Address
	TokenAmount   big.Int
	TokenDecimals uint8
	Token         common.Address
	Time          uint64
	BlockNumber   uint64
	TxHash        common.Hash
}

type LockLogsIndexer interface {
	Address() []common.Address
	IndexLogs(ctx context.Context, logs []types.Log) ([]Lock, error)
	ProcessLog(ctx context.Context, vLog types.Log) (*Lock, error)
}

type LogIndexer[T any] interface {
	Address() []common.Address
	Index(context.Context, ...types.Log) ([]T, error)
}

type BlockPprocessor[T any] interface {
	Process(context.Context, ...types.Block) (T, error)
}

func GetLogs(ctx context.Context, client *ethclient.Client, addresses []common.Address, blocknumber int64) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		Addresses: addresses,
		FromBlock: big.NewInt(blocknumber),
		ToBlock:   big.NewInt(blocknumber),
	}

	logs, err := client.FilterLogs(ctx, query)
	if err != nil {
		return nil, err
	}

	return logs, nil
}
