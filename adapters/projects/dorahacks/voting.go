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
	// https://taikoscan.io/address/0x6c394B04C35A51aB2CafE9dd8511515A898cb419
	VotingAddress string = "0x6c394B04C35A51aB2CafE9dd8511515A898cb419"

	// event PublishMessage(uint256 indexed _msgIdx, Message _message, PubKey _encPubKey)
	logPublishMessageSignatureHex string = "0x8bb5a8cf78a5b2f53c73e2feacb1fb3e91c3f03cb15e33f53174db20e37e3928"
)

type VotingIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewVotingIndexer(client *ethclient.Client, addresses []common.Address) *VotingIndexer {
	return &VotingIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &VotingIndexer{}

func (indexer *VotingIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *VotingIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var ws []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isPublishMessage(l) {
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
				TxHash:      txHash,
			},
			User: sender, // msg.sender equals voter for DoraHacks MACI
		}

		ws = append(ws, *w)
	}

	return ws, nil
}

func (indexer *VotingIndexer) isPublishMessage(l types.Log) bool {
	return l.Topics[0].Hex() == logPublishMessageSignatureHex
}
