package adapters

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// ProcessBlockType is a function type that processes blocks for a given chain ID.
type ProcessBlockType func(ctx context.Context, chainID *big.Int, block *types.Block) error

// ProcessFunc is an implementation of ProcessFuncType that processes blocks.
func ProcessBlockFunc(ctx context.Context, chainID *big.Int, block *types.Block) error {
	if ctx == nil {
		return errors.New("context is nil")
	}
	if chainID == nil {
		return errors.New("chainID is nil")
	}
	if block == nil {
		return errors.New("block is nil")
	}

	// Process the block. This is a placeholder for actual processing logic.
	processBlock(ctx, chainID, block)

	return nil
}

func processBlock(ctx context.Context, chainID *big.Int, block *types.Block) {
	// Add your block processing logic here.
	// This is just a placeholder for demonstration purposes.
	// txs := block.Transactions()
	// for _, _ := range txs {

	// }
}
