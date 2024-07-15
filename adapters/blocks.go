package adapters

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

// BlockProcessor is an interface that defines the methods for processing blocks.
type BlockProcessor interface {
	ProcessBlock(ctx context.Context, block *types.Block, client *ethclient.Client) (*[]common.Address, error)
}
