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

type Whitelist struct {
	User        common.Address
	Time        uint64
	BlockNumber uint64
	TxHash      common.Hash
}

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

type LogIndexer[T any] interface {
	Addresses() []common.Address
	Index(context.Context, ...types.Log) ([]T, error)
}

type Prediction struct {
	User          common.Address
	TokenAmount   *big.Int
	TokenDecimals uint8
	Token         common.Address
	EventCode     string
	WinnerIndex   uint8 // 0 or 1
	BlockTime     uint64
	BlockNumber   uint64
	TxHash        common.Hash
}

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
