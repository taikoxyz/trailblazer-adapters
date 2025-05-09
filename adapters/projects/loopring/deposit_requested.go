package loopring

import (
	"context"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/taikoxyz/trailblazer-adapters/adapters"
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/loopring_exchangev3"
)

const (
	// https://taikoscan.io/address/0xbD787F374198d29E2F8Fa228c778FE39e1a5d3a9
	DepositRequestedAddress string = "0xbD787F374198d29E2F8Fa228c778FE39e1a5d3a9"

	logDepositRequestedSignature string = "DepositRequested(address,address,address,uint16,uint96)"
)

type DepositRequestedIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewDepositRequestedIndexer(client *ethclient.Client, addresses []common.Address) *DepositRequestedIndexer {
	return &DepositRequestedIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &DepositRequestedIndexer{}

func (indexer *DepositRequestedIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *DepositRequestedIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var ws []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isDepositRequested(l) {
			continue
		}

		var depositRequestedEvent struct {
			From    common.Address
			To      common.Address
			Token   common.Address
			TokenId uint16
			Amount  *big.Int
		}

		loopringABI, err := abi.JSON(strings.NewReader(loopring_exchangev3.LoopringExchangev3ABI))
		if err != nil {
			return nil, err
		}

		err = loopringABI.UnpackIntoInterface(&depositRequestedEvent, "DepositRequested", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			User:        depositRequestedEvent.From,
			Time:        block.Time(),
			BlockNumber: block.NumberU64(),
			TxHash:      l.TxHash,
		}

		ws = append(ws, *w)
	}

	return ws, nil
}

func (indexer *DepositRequestedIndexer) isDepositRequested(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logDepositRequestedSignature)).Hex()
}
