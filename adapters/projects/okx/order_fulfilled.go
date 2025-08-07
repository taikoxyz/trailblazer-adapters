package okx

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/order"
)

const (
	// https://taikoscan.io/address/0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC
	OrderFulfilledAddress string = "0x00000000000000ADc04C56Bf30aC9d3c0aAF14dC"

	logOrderFulfilledSignature string = "OrderFulfilled(bytes32,address,address,address,(uint8,address,uint256,uint256)[],(uint8,address,uint256,uint256,address)[])"
)

type SpentItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
}

type ReceivedItem struct {
	ItemType   uint8
	Token      common.Address
	Identifier *big.Int
	Amount     *big.Int
	Recipient  common.Address
}

type OrderFulfilledEvent struct {
	OrderHash     [32]byte
	Offerer       common.Address
	Zone          common.Address
	Recipient     common.Address
	Offer         []SpentItem
	Consideration []ReceivedItem
}

type OrderFulfilledIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewOrderFulfilledIndexer(client *ethclient.Client, addresses []common.Address) *OrderFulfilledIndexer {
	return &OrderFulfilledIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &OrderFulfilledIndexer{}

func (indexer *OrderFulfilledIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *OrderFulfilledIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isOrderFulfilledLog(l) {
			continue
		}

		var orderFulfilledEvent OrderFulfilledEvent

		orderFulfilledABI, err := abi.JSON(strings.NewReader(order.ABI))
		if err != nil {
			return nil, err
		}
		err = orderFulfilledABI.UnpackIntoInterface(&orderFulfilledEvent, "OrderFulfilled", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			Metadata: adapters.Metadata{
				BlockTime:   block.Time(),
				BlockNumber: block.NumberU64(),
				TxHash:      l.TxHash,
			},
			User: orderFulfilledEvent.Recipient,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func (indexer *OrderFulfilledIndexer) isOrderFulfilledLog(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logOrderFulfilledSignature)).Hex()
}
