package loopex

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/sale"
)

const (
	// https://taikoscan.io/address/0x8733764434c3a3C2b2Fa3A5033CA49FDdDF976C0
	NewSaleAddress string = "0x8733764434c3a3C2b2Fa3A5033CA49FDdDF976C0"

	logNewSaleSignature string = "NewSale(address,uint256,address,uint256,address,uint256,uint256)"
)

type NewSaleEvent struct {
	Seller    common.Address `json:"seller"`
	TokenId   *big.Int       `json:"tokenId"`
	Buyer     common.Address `json:"buyer"`
	Price     *big.Int       `json:"price"`
	Currency  common.Address `json:"currency"`
	Amount    *big.Int       `json:"amount"`
	Timestamp *big.Int       `json:"timestamp"`
}

type NewSaleIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewNewSaleIndexer(client *ethclient.Client, addresses []common.Address) *NewSaleIndexer {
	return &NewSaleIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &NewSaleIndexer{}

func (indexer *NewSaleIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *NewSaleIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isNewSaleLog(l) {
			continue
		}

		var newSaleEvent NewSaleEvent

		newSaleABI, err := abi.JSON(strings.NewReader(sale.ABI))
		if err != nil {
			return nil, err
		}

		err = newSaleABI.UnpackIntoInterface(&newSaleEvent, "NewSale", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			User:        newSaleEvent.Buyer,
			Time:        block.Time(),
			BlockNumber: block.NumberU64(),
			TxHash:      l.TxHash,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func (indexer *NewSaleIndexer) isNewSaleLog(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logNewSaleSignature)).Hex()
}
