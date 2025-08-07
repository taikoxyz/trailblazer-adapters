package dorahacks

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
)

const (
	// event SignUp(uint256 indexed _stateIdx, PubKey _userPubKey, uint256 _voiceCreditBalance);
	logSignupSignatureHex string = "0xc7563c66f89e2fb0839e2b64ed54fe4803ff9428777814772ccfe4c385072c4b"
)

type SignupIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewSignupIndexer(client *ethclient.Client, addresses []common.Address) *SignupIndexer {
	return &SignupIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &SignupIndexer{}

func (indexer *SignupIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *SignupIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var ws []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isSignup(l) {
			continue
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		txHash := l.TxHash
		receipt, err := indexer.client.TransactionReceipt(ctx, txHash)
		if err != nil {
			return nil, err
		}
		if receipt.Status == 0 {
			// tx failed, continue
			continue
		}

		tx, pending, err := indexer.client.TransactionByHash(ctx, txHash)
		if err != nil {
			return nil, err
		}
		if pending {
			// tx pending, continue
			continue
		}

		sender, err := indexer.client.TransactionSender(ctx, tx, block.Hash(), receipt.TransactionIndex)
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			Metadata: adapters.Metadata{
				BlockTime:   block.Time(),
				BlockNumber: block.NumberU64(),
				TxHash:      l.TxHash,
			},
			User: sender, // msg.sender equals signup for DoraHacks MACI
		}

		ws = append(ws, *w)
	}

	return ws, nil
}

func (indexer *SignupIndexer) isSignup(l types.Log) bool {
	return l.Topics[0].Hex() == logSignupSignatureHex
}
