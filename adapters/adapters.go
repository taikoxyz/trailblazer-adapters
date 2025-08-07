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

type Metadata struct {
	BlockTime   uint64
	BlockNumber uint64
	TxHash      common.Hash
}

// Whitelist contains information for general whitelisted protocols
type Whitelist struct {
	Metadata
	User common.Address
}

// LPTransfer is used for LP transfers.
// For examples, see Izumi and Ritsu.
type LPTransfer struct {
	Metadata
	From           common.Address
	To             common.Address
	Token0Amount   *big.Int
	Token0Decimals uint8
	Token0         common.Address
	Token1Amount   *big.Int
	Token1Decimals uint8
	Token1         common.Address
}

// Lock is used for token locking campaign.
// For examples, see Drips and Symmetric.
type Lock struct {
	Metadata
	User          common.Address
	TokenAmount   *big.Int
	TokenDecimals uint8
	Token         common.Address
	Duration      uint64
}

// Position is used for airdrop or trading campaign.
// For examples, see Avalon.
type Position struct {
	Metadata
	User          common.Address
	TokenAmount   *big.Int
	TokenDecimals uint8
	Token         common.Address
}

// Prediction is used for prediction campaign.
// For examples, see Robinos.
type Prediction struct {
	Metadata
	User          common.Address
	TokenAmount   *big.Int
	TokenDecimals uint8
	Token         common.Address
	EventCode     string
}

type Nft struct {
	Collection common.Address
	TokenId    *big.Int
}

// NftSold is used for NFT sold events.
// For examples, see Okidori marketplace.
type NftSold struct {
	Metadata
	Nft
	Seller   common.Address
	Buyer    common.Address
	Price    *big.Int
	Currency common.Address
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
