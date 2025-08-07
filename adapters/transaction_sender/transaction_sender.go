package transactionsender

import (
	"context"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

// https://taikoscan.io/address/0x1Df2De291F909baA50C1456C87C71Edf9Fb199D5
const Recipient string = "0x1Df2De291F909baA50C1456C87C71Edf9Fb199D5"

type Processor struct {
	client          *ethclient.Client
	ValidRecipients map[string]struct{}
}

func New(client *ethclient.Client) *Processor {
	return &Processor{
		client: client,
		ValidRecipients: map[string]struct{}{
			Recipient: {},
		},
	}
}

var _ adapters.BlockProcessor[adapters.Whitelist] = &Processor{}

func (processor *Processor) Process(ctx context.Context, blocks ...*types.Block) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, b := range blocks {
		txs := b.Transactions()

		for _, tx := range txs {
			to := tx.To()
			if to == nil {
				continue
			}

			receipt, err := processor.client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				return nil, err
			}
			if receipt.Status == 0 {
				continue
			}

			sender, err := processor.client.TransactionSender(ctx, tx, b.Hash(), receipt.TransactionIndex)
			if err != nil {
				return nil, err
			}

			if _, exists := processor.ValidRecipients[to.Hex()]; exists {
				whitelist = append(whitelist, adapters.Whitelist{
					Metadata: adapters.Metadata{
						BlockTime:   b.Time(),
						BlockNumber: b.NumberU64(),
						TxHash:      tx.Hash(),
					},
					User: sender,
				})
			}
		}
	}

	return whitelist, nil
}
