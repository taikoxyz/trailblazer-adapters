package blocks

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

var _ adapters.BlockProcessor = (*TransactionSender)(nil)

type TransactionSender struct {
	ValidRecipient map[string]struct{}
}

// NewTransactionSender creates a new NewTransactionSender.
func NewTransactionSender() *TransactionSender {
	return &TransactionSender{ValidRecipient: map[string]struct{}{
		common.HexToAddress("0x1Df2De291F909baA50C1456C87C71Edf9Fb199D5").Hex(): {},
	}}
}

func (indexer *TransactionSender) ProcessBlock(ctx context.Context, block *types.Block, client *ethclient.Client) (*[]common.Address, error) {
	txs := block.Transactions()
	var result []common.Address

	for _, tx := range txs {
		receipt, err := client.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, err
		}
		sender, err := client.TransactionSender(ctx, tx, block.Hash(), receipt.TransactionIndex)
		if err != nil {
			return nil, err
		}
		to := tx.To()
		if to == nil {
			continue
		}
		if _, exists := indexer.ValidRecipient[to.Hex()]; exists {
			result = append(result, sender)
		}
	}
	return &result, nil
}
