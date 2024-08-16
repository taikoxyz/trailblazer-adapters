package adapters

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

// BlockProcessor is an interface that defines the methods for processing blocks.
type BlockProcessor interface {
	ProcessBlock(ctx context.Context, block *types.Block, client *ethclient.Client) ([]Whitelist, error)
}

type Whitelist struct {
	User        common.Address
	Time        uint64
	BlockNumber uint64
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
}
