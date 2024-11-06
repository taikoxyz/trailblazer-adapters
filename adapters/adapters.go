package adapters

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// https://taikoscan.io/address/0xA51894664A773981C6C112C43ce576f315d5b1B6
	WETHAddress string = "0xA51894664A773981C6C112C43ce576f315d5b1B6"
	// https://taikoscan.io/address/0xa9d23408b9ba935c230493c40c73824df71a0975
	TaikoTokenAddress  string = "0xa9d23408b9ba935c230493c40c73824df71a0975"
	TaikoTokenDecimals uint8  = 18
)

func ZeroAddress() common.Address {
	return common.HexToAddress("0x0000000000000000000000000000000000000000")
}

// Whitelist contains information for general whitelisted protocols
type Whitelist struct {
	User        common.Address
	Time        uint64
	BlockNumber uint64
	TxHash      common.Hash
}

// LPTransfer is used for LP transfers.
// For examples, see Izumi and Ritsu.
type LPTransfer struct {
	From           common.Address
	To             common.Address
	Token0Amount   *big.Int
	Token0Decimals uint8
	Token0         common.Address
	Token1Amount   *big.Int
	Token1Decimals uint8
	Token1         common.Address
	Time           uint64
	BlockNumber    uint64
	TxHash         common.Hash
}

// Lock is used for token locking campaign.
// For examples, see Drips and Symmetric.
type Lock struct {
	User          common.Address
	TokenAmount   *big.Int
	TokenDecimals uint8
	Token         common.Address
	Duration      uint64
	BlockTime     uint64
	BlockNumber   uint64
	TxHash        common.Hash
}

// Position is used for logx trading campaign
type Position struct {
	User          common.Address
	TokenAmount   *big.Int
	TokenDecimals uint8
	Token         common.Address
	BlockTime     uint64
	BlockNumber   uint64
	TxHash        common.Hash
}

// Prdiction is used for prediction campaign.
// For examples, see Robinos.
type Prediction struct {
	User          common.Address
	TokenAmount   *big.Int
	TokenDecimals uint8
	Token         common.Address
	EventCode     string
	BlockTime     uint64
	BlockNumber   uint64
	TxHash        common.Hash
}

// LogIndexer is the generic interface for indexing logs into types above.
// By indexing, it parses information of logs.
type LogIndexer[T any] interface {
	// Addresses returns indexed contract addresses.
	Addresses() []common.Address
	// Index transforms given logs into types above.
	Index(context.Context, ...types.Log) ([]T, error)
}

// BlockProcessor is the generic interface for indexing block information.
// By block processing, it parses information of blocks.
type BlockProcessor[T any] interface {
	Process(context.Context, ...*types.Block) ([]T, error)
}

// GetLogs returns []types.Log for given addresses and blocknumber
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

// GetBlock return *types.Block by blocknumber
func GetBlock(ctx context.Context, client *ethclient.Client, blocknumber int64) (*types.Block, error) {
	block, err := client.BlockByNumber(context.Background(), big.NewInt(blocknumber))
	if err != nil {
		return nil, err
	}

	return block, nil
}
