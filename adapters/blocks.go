package adapters

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockProcessor is an interface that defines the methods for processing blocks.
type BlockProcessor interface {
	ProcessBlock(ctx context.Context, block *types.Block, client *ethclient.Client) ([]Whitelist, error)
}
