package conft

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
	"github.com/taikoxyz/trailblazer-adapters/adapters/contracts/conft"
)

const (
	// https://taikoscan.io/address/0x6Ce2CFD7674cf47A851690a11d1DB45c6cCBe17A
	TokenSoldAddress string = "0x6Ce2CFD7674cf47A851690a11d1DB45c6cCBe17A"

	logTokenSoldSignature string = "TokenSold(uint128,address,address,address,uint256,uint128)"
)

type TokenSoldEvent struct {
	Id              *big.Int
	Seller          common.Address
	Buyer           common.Address
	ContractAddress common.Address
	TokenId         *big.Int
	Price           *big.Int
}

type TokenSoldIndexer struct {
	client    *ethclient.Client
	addresses []common.Address
}

func NewTokenSoldIndexer(client *ethclient.Client, addresses []common.Address) *TokenSoldIndexer {
	return &TokenSoldIndexer{
		client:    client,
		addresses: addresses,
	}
}

var _ adapters.LogIndexer[adapters.Whitelist] = &TokenSoldIndexer{}

func (indexer *TokenSoldIndexer) Addresses() []common.Address {
	return indexer.addresses
}

func (indexer *TokenSoldIndexer) Index(ctx context.Context, logs ...types.Log) ([]adapters.Whitelist, error) {
	var whitelist []adapters.Whitelist

	for _, l := range logs {
		if !indexer.isTokenSoldLog(l) {
			continue
		}

		var tokenSoldEvent TokenSoldEvent

		tokenSoldABI, err := abi.JSON(strings.NewReader(conft.ABI))
		if err != nil {
			return nil, err
		}

		err = tokenSoldABI.UnpackIntoInterface(&tokenSoldEvent, "TokenSold", l.Data)
		if err != nil {
			return nil, err
		}

		block, err := indexer.client.BlockByNumber(ctx, big.NewInt(int64(l.BlockNumber)))
		if err != nil {
			return nil, err
		}

		w := &adapters.Whitelist{
			User:        tokenSoldEvent.Buyer,
			Time:        block.Time(),
			BlockNumber: block.NumberU64(),
			TxHash:      l.TxHash,
		}

		whitelist = append(whitelist, *w)
	}

	return whitelist, nil
}

func (indexer *TokenSoldIndexer) isTokenSoldLog(l types.Log) bool {
	return l.Topics[0].Hex() == crypto.Keccak256Hash([]byte(logTokenSoldSignature)).Hex()
}
